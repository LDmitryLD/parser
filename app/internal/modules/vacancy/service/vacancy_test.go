package service

import (
	"fmt"
	"projects/LDmitryLD/parser/app/internal/models"
	smocks "projects/LDmitryLD/parser/app/internal/modules/vacancy/storage/mocks"
	pmocks "projects/LDmitryLD/parser/app/internal/parser/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	query   = "golang"
	id      = 1
	vacancy = models.Vacancy{
		Context: "context",
		Type:    "type",
	}
	vacs = []models.Vacancy{vacancy}
)

func TestVacancyService_Search(t *testing.T) {
	storageMock := smocks.NewVacancyStorager(t)
	storageMock.On("Search", query).Return([]models.Vacancy{vacancy}, nil)

	vacService := VacancyService{
		storage: storageMock,
	}

	got, err := vacService.Search(query)

	assert.Equal(t, []models.Vacancy{vacancy}, got)
	assert.Nil(t, err)
}

func TestVacancyService_Search_Parser(t *testing.T) {
	storageMock := smocks.NewVacancyStorager(t)
	storageMock.On("Search", query).Return([]models.Vacancy{}, fmt.Errorf("test error"))
	storageMock.On("Create", vacs)

	parserMock := pmocks.NewParser(t)
	parserMock.On("Search", query).Return(vacs, nil)

	vacService := NewVacanceService(parserMock, storageMock)

	got, err := vacService.Search(query)

	assert.Equal(t, vacs, got)
	assert.Nil(t, err)
}

func TestVacancyService_Search_Error(t *testing.T) {
	storageMock := smocks.NewVacancyStorager(t)
	storageMock.On("Search", query).Return([]models.Vacancy{}, fmt.Errorf("test error"))

	parserMock := pmocks.NewParser(t)
	parserMock.On("Search", query).Return([]models.Vacancy{}, fmt.Errorf("test error"))

	vacSerice := NewVacanceService(parserMock, storageMock)

	_, err := vacSerice.Search(query)

	assert.NotNil(t, err)
}

func TestVacancyService_Get(t *testing.T) {
	storageMock := smocks.NewVacancyStorager(t)
	storageMock.On("SearchByID", id).Return(vacancy, nil)

	vacService := VacancyService{
		storage: storageMock,
	}

	got, err := vacService.Get(id)

	assert.Equal(t, vacancy, got)
	assert.Nil(t, err)
}

func TestVacancyService_List(t *testing.T) {
	storageMock := smocks.NewVacancyStorager(t)
	storageMock.On("List").Return(vacs, nil)

	vacService := VacancyService{
		storage: storageMock,
	}

	got, err := vacService.List()

	assert.Equal(t, vacs, got)
	assert.Nil(t, err)
}

func TestVacancyService_Delete(t *testing.T) {
	storageMock := smocks.NewVacancyStorager(t)
	storageMock.On("Delete", id).Return(nil)

	vacService := VacancyService{
		storage: storageMock,
	}

	err := vacService.Delete(id)

	assert.Nil(t, err)
}
