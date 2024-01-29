package storage

import (
	"fmt"
	"projects/LDmitryLD/parser/app/internal/db/adapter/mocks"
	"projects/LDmitryLD/parser/app/internal/infrastructure/errors"
	"projects/LDmitryLD/parser/app/internal/models"
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

func TestVacancyStorage_Search(t *testing.T) {
	adapterMock := mocks.NewSQLAdapterer(t)
	adapterMock.On("SelectVacancies", query, false).Return(vacs, nil)

	vacStorage := NewVacancyStorage(adapterMock)

	got, err := vacStorage.Search(query)

	assert.Equal(t, vacs, got)
	assert.Nil(t, err)
}

func TestVacancyStorage_SearchByID(t *testing.T) {
	adapterMock := mocks.NewSQLAdapterer(t)
	adapterMock.On("SelectVacancy", id).Return(vacancy, nil)

	vacStorage := NewVacancyStorage(adapterMock)

	vacStorage.vacsByID[id] = true

	got, err := vacStorage.SearchByID(id)

	assert.Equal(t, vacancy, got)
	assert.Nil(t, err)
}

func TestVacancyStorage_SearchByID_NotFound(t *testing.T) {
	adapterMock := mocks.NewSQLAdapterer(t)

	vacStorage := NewVacancyStorage(adapterMock)

	_, err := vacStorage.SearchByID(id)

	assert.NotNil(t, err)
}

func TestVacancyStorage_SearchByID_Error(t *testing.T) {
	adapterMock := mocks.NewSQLAdapterer(t)
	adapterMock.On("SelectVacancy", id).Return(models.Vacancy{}, fmt.Errorf("test error"))

	vacStorage := NewVacancyStorage(adapterMock)

	vacStorage.vacsByID[id] = true

	_, err := vacStorage.SearchByID(id)

	assert.NotNil(t, err)
}

func TestVacancyStorage_List(t *testing.T) {
	adapterMock := mocks.NewSQLAdapterer(t)
	adapterMock.On("SelectVacancies", "", true).Return(vacs, nil)

	vacStorage := NewVacancyStorage(adapterMock)

	got, err := vacStorage.List()

	assert.Equal(t, vacs, got)
	assert.Nil(t, err)
}

func TestVacancyStorage_Create(t *testing.T) {
	adapterMock := mocks.NewSQLAdapterer(t)
	adapterMock.On("Insert", vacancy).Return(id, nil)

	vacStorage := NewVacancyStorage(adapterMock)

	vacStorage.Create(vacs)

	got := vacStorage.vacsByID[id]

	assert.True(t, got)
}

func TestVacancyStorage_Delete(t *testing.T) {
	adaoterMock := mocks.NewSQLAdapterer(t)
	adaoterMock.On("Delete", id).Return(nil)

	vacStorage := NewVacancyStorage(adaoterMock)

	vacStorage.vacsByID[id] = true

	err := vacStorage.Delete(id)

	assert.Nil(t, err)
	assert.False(t, vacStorage.vacsByID[id])
}

func TestVacancyStorage_Delete_NotFound(t *testing.T) {
	adaoterMock := mocks.NewSQLAdapterer(t)

	vacStorage := NewVacancyStorage(adaoterMock)

	err := vacStorage.Delete(id)

	assert.Equal(t, errors.ErrNotFound, err)
}

func TestVacancyStorage_Delete_Error(t *testing.T) {
	adaoterMock := mocks.NewSQLAdapterer(t)
	adaoterMock.On("Delete", id).Return(fmt.Errorf("test error"))

	vacStorage := NewVacancyStorage(adaoterMock)

	vacStorage.vacsByID[id] = true

	err := vacStorage.Delete(id)

	assert.NotNil(t, err)
	assert.True(t, vacStorage.vacsByID[id])
}
