package storage

import (
	"fmt"
	"log"
	"projects/LDmitryLD/parser/app/internal/db/adapter"
	"projects/LDmitryLD/parser/app/internal/infrastructure/errors"
	"projects/LDmitryLD/parser/app/internal/models"
	"sync"
)

type VacancyStorager interface {
	Search(query string) ([]models.Vacancy, error)
	SearchByID(id int) (models.Vacancy, error)
	List() ([]models.Vacancy, error)
	Create(vacs []models.Vacancy)
	Delete(id int) error
}

type VacancyStorage struct {
	adapter  adapter.SQLAdapterer
	vacsByID map[int]bool
	sync.Mutex
}

func NewVacancyStorage(adapter adapter.SQLAdapterer) *VacancyStorage {
	return &VacancyStorage{
		adapter:  adapter,
		vacsByID: make(map[int]bool, 25),
	}
}

func (v *VacancyStorage) Search(query string) ([]models.Vacancy, error) {
	return v.adapter.SelectVacancies(query, false)
}

func (v *VacancyStorage) SearchByID(id int) (models.Vacancy, error) {
	v.Lock()
	defer v.Unlock()

	if _, ok := v.vacsByID[id]; !ok {
		log.Printf("вакансии с id % d нет в базе", id)
		log.Println(v.vacsByID)
		return models.Vacancy{}, errors.ErrNotFound
	}

	log.Printf("вакансии с id % d eсть в базе, сейчас найдём", id)
	log.Println(v.vacsByID)

	vacancy, err := v.adapter.SelectVacancy(id)
	if err != nil {
		log.Printf("вакансия с id %d не найдена, ошибка: %s\n", id, err.Error())
		return models.Vacancy{}, err
	}

	v.vacsByID[id] = true

	return vacancy, nil
}

func (v *VacancyStorage) List() ([]models.Vacancy, error) {
	return v.adapter.SelectVacancies("", true)
}

func (v *VacancyStorage) Create(vacs []models.Vacancy) {
	for _, vac := range vacs {
		v.Lock()
		id, err := v.adapter.Insert(vac)
		fmt.Println("ID:", id)
		if err != nil {
			log.Println("ошибка при добавлении вакансии в бд:", err)
			continue
		}

		log.Println("Добавлено в кэш")

		v.vacsByID[id] = true
		v.Unlock()
	}
}

func (v *VacancyStorage) Delete(id int) error {
	v.Lock()
	defer v.Unlock()

	if _, ok := v.vacsByID[id]; !ok {
		return errors.ErrNotFound
	}

	if err := v.adapter.Delete(id); err != nil {
		log.Println("ошибка при удалнии вакансии")
		return nil
	}

	delete(v.vacsByID, id)

	return nil
}
