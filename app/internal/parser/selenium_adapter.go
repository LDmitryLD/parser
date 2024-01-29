package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"projects/LDmitryLD/parser/app/internal/models"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/tebeka/selenium"
)

const (
	habrLink = "https://career.habr.com/"
	timeout  = 50 * time.Second
)

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name=Parser
type Parser interface {
	Search(query string) ([]models.Vacancy, error)
}

type SeleniumParser struct {
	wd selenium.WebDriver
}

func NewParser() Parser {

	wd := NewWebDriver()
	return &SeleniumParser{
		wd: wd,
	}
}

func (p *SeleniumParser) Search(query string) ([]models.Vacancy, error) {
	if err := p.getPage(1, query); err != nil {
		return nil, err
	}

	pagesCount, err := p.pagesCount()
	if err != nil {
		return nil, err
	}

	log.Println("Pages count:", pagesCount)

	var mu sync.Mutex
	var wg sync.WaitGroup
	var vacs []models.Vacancy

	for i := 1; i <= pagesCount; i++ {
		p.getPage(i, query)
		links, err := p.getLinks()
		if err != nil {
			log.Println("ошибка при полученнии списка вакансий: ", err)
			return nil, err
		}
		for _, link := range links {
			wg.Add(1)
			go func(vacLink string) {
				defer wg.Done()

				vacancyRaw, err := getVacancy(vacLink)
				if err != nil {
					log.Println("ошибка при получении ссылки на вакансию: ", err)
				}
				var vac models.Vacancy
				if err := json.Unmarshal([]byte(vacancyRaw), &vac); err != nil {
					log.Println("ошибка при анмаршалинге вакансии:", err)
					log.Println("Вакансия с которой произошла ошибка: ", vacancyRaw)
				}
				mu.Lock()
				vacs = append(vacs, vac)
				mu.Unlock()
			}(link)
		}
	}
	wg.Wait()

	fmt.Println("Получено ваканский: ", len(vacs))

	return vacs, nil
}

func getVacancy(vacLink string) (string, error) {
	resp, err := http.Get(vacLink)
	if err != nil {
		log.Printf("ошибка при http.Get(%s): %s\n", vacLink, err.Error())
		return "", err
	}

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("ошибка при получении докумена из запроса на вакансию:", err)
		return "", err
	}

	dd := doc.Find("script[type=\"application/ld+json\"]")
	if dd == nil {
		log.Println("habr vacancy nodes not found")
		return "", err
	}

	ss := dd.First().Text()

	return ss, nil
}

func (p *SeleniumParser) getPage(pageNum int, query string) error {
	err := p.wd.Get(fmt.Sprintf("https://career.habr.com/vacancies?page=%d&q=%s&type=all", pageNum, query))
	if err != nil {
		log.Println("ошибка при получении первой страницы", err)
		return err
	}
	time.Sleep(timeout)

	return nil
}

func (p *SeleniumParser) getLinks() ([]string, error) {
	elems, err := p.wd.FindElements(selenium.ByCSSSelector, ".vacancy-card__title-link")
	if err != nil {
		return nil, err
	}

	var links []string
	for i := range elems {
		var link string
		link, err := elems[i].GetAttribute("href")
		if err != nil {
			continue
		}
		links = append(links, habrLink+link)
	}

	return links, nil
}

func (p *SeleniumParser) pagesCount() (int, error) {
	elem, err := p.wd.FindElement(selenium.ByCSSSelector, ".search-total")
	if err != nil {
		log.Println("ошибка при получении общего кол-ва вакансий:", err)
		return 0, err
	}

	vacancyCountRaw, err := elem.Text()
	log.Println("vacancyCountRaw::", vacancyCountRaw)
	if err != nil {
		log.Println("ошибка при получении кол-ва вакансий:", err)
		return 0, nil
	}

	return pagesCount(vacancyCountRaw)
}

func pagesCount(countRaw string) (int, error) {
	// count, err := strconv.Atoi(countRaw)
	// if err != nil {
	// 	log.Println("ошибка при приведении countRaw к int", err)
	// 	return 0, err
	// }

	numberStr := strings.TrimPrefix(countRaw, "Найдено ")
	numberStr = strings.TrimSuffix(numberStr, " вакансий")
	count, err := strconv.Atoi(numberStr)
	if err != nil {
		log.Println("Ошибка при преобразовании строки в число:", err)
		return 0, err
	}

	pageCount := math.Ceil(float64(count) / 25.0)

	return int(pageCount), nil
}
