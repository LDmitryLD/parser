package service

import (
	"log"
	"projects/LDmitryLD/parser/app/internal/models"
	"projects/LDmitryLD/parser/app/internal/modules/vacancy/storage"
	"projects/LDmitryLD/parser/app/internal/parser"
)

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name=Vacancyer
type Vacancyer interface {
	Search(query string) ([]models.Vacancy, error)
	Get(id int) (models.Vacancy, error)
	List() ([]models.Vacancy, error)
	Delete(id int) error
}

type VacancyService struct {
	storage storage.VacancyStorager
	parser  parser.Parser
}

func NewVacanceService(parser parser.Parser, storage storage.VacancyStorager) *VacancyService {
	return &VacancyService{
		storage: storage,
		parser:  parser,
	}
}

func (v *VacancyService) Search(query string) ([]models.Vacancy, error) {
	var vacs []models.Vacancy
	var err error

	vacs, err = v.storage.Search(query)
	if err == nil {
		log.Println("вакансии получены и БД")
		return vacs, nil
	}

	vacs, err = v.parser.Search(query)
	if err != nil {
		log.Println("ошибка при поиске ваканский:", err)
		return nil, err
	}

	v.storage.Create(vacs)

	return vacs, nil
}

func (v *VacancyService) Get(id int) (models.Vacancy, error) {
	return v.storage.SearchByID(id)
}

func (v *VacancyService) List() ([]models.Vacancy, error) {
	return v.storage.List()
}

func (v *VacancyService) Delete(id int) error {
	return v.storage.Delete(id)
}
