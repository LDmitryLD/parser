package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"projects/LDmitryLD/parser/app/internal/infrastructure/errors"
	"projects/LDmitryLD/parser/app/internal/models"
	"projects/LDmitryLD/parser/app/internal/modules/vacancy/service/mocks"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

var (
	query   = "golang"
	id      = 1
	vacancy = models.Vacancy{
		Context: "context",
		Type:    "type",
	}
)

func TestVacancyController_Search_BadRequest(t *testing.T) {
	req := map[string]interface{}{"query": 1}
	reqJSON, _ := json.Marshal(req)

	serviceMock := mocks.NewVacancyer(t)

	vacController := NewVacancyController(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(vacController.Search))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestVacancyController_Search_Internal(t *testing.T) {
	expect := models.ApiResponse{
		Code:    http.StatusInternalServerError,
		Message: "test error",
	}
	req := SearchRequest{
		Query: query,
	}
	reqJSON, _ := json.Marshal(req)

	serviceMock := mocks.NewVacancyer(t)
	serviceMock.On("Search", query).Return([]models.Vacancy{}, fmt.Errorf("test error"))

	vacController := NewVacancyController(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(vacController.Search))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	var respSearch models.ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&respSearch); err != nil {
		t.Fatal(errors.ErrTestDecode, err)
	}

	assert.Equal(t, expect, respSearch)
}

func TestVacancyController_Search(t *testing.T) {
	expect := []models.Vacancy{vacancy}
	req := SearchRequest{
		Query: query,
	}
	reqJSON, _ := json.Marshal(req)

	serviceMock := mocks.NewVacancyer(t)
	serviceMock.On("Search", query).Return(expect, nil)

	vacController := NewVacancyController(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(vacController.Search))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(errors.ErrTestReq, err)
	}
	defer resp.Body.Close()

	var respSearch []models.Vacancy
	if err := json.NewDecoder(resp.Body).Decode(&respSearch); err != nil {
		t.Fatal(errors.ErrTestDecode, err)
	}

	assert.Equal(t, expect, respSearch)
}

func TestVacancyController_Delete_BadRequest(t *testing.T) {
	serviceMock := mocks.NewVacancyer(t)

	vacController := NewVacancyController(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/parser/delete/i", nil)

	r := chi.NewRouter()
	r.MethodFunc("DELETE", "/parser/delete/{id}", vacController.Delete)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestVacancyController_Delete_Internal(t *testing.T) {
	serviceMock := mocks.NewVacancyer(t)
	serviceMock.On("Delete", id).Return(fmt.Errorf("test error"))

	vacController := NewVacancyController(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/parser/delete/1", nil)

	r := chi.NewRouter()
	r.MethodFunc("DELETE", "/parser/delete/{id}", vacController.Delete)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestVacancyController_Delete_NotFound(t *testing.T) {
	serviceMock := mocks.NewVacancyer(t)
	serviceMock.On("Delete", id).Return(errors.ErrNotFound)

	vacController := NewVacancyController(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/parser/delete/1", nil)

	r := chi.NewRouter()
	r.MethodFunc("DELETE", "/parser/delete/{id}", vacController.Delete)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestVacancyController_Delete(t *testing.T) {
	expect := models.ApiResponse{
		Code:    200,
		Message: fmt.Sprintf("вакансия с id %d удалена", id),
	}

	serviceMock := mocks.NewVacancyer(t)
	serviceMock.On("Delete", id).Return(nil)

	vacController := NewVacancyController(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/parser/delete/1", nil)

	r := chi.NewRouter()
	r.MethodFunc("DELETE", "/parser/delete/{id}", vacController.Delete)

	r.ServeHTTP(rr, req)

	var resp models.ApiResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(errors.ErrTestDecode)
	}

	assert.Equal(t, expect, resp)
}

func TestVacancyController_Get_BadRequest(t *testing.T) {
	serviceMock := mocks.NewVacancyer(t)

	vacController := NewVacancyController(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/parser/delete/i", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/parser/delete/{id}", vacController.Get)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestVacancyCotroller_Get_Internal(t *testing.T) {
	serviceMock := mocks.NewVacancyer(t)
	serviceMock.On("Get", id).Return(models.Vacancy{}, fmt.Errorf("test error"))

	vacController := NewVacancyController(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/parser/1", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/parser/{id}", vacController.Get)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestVacancyController_Get_NotFound(t *testing.T) {
	serviceMock := mocks.NewVacancyer(t)
	serviceMock.On("Get", id).Return(models.Vacancy{}, errors.ErrNotFound)

	vacController := NewVacancyController(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/parser/1", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/parser/{id}", vacController.Get)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestVacancyController_Get(t *testing.T) {
	serviceMock := mocks.NewVacancyer(t)
	serviceMock.On("Get", id).Return(vacancy, nil)

	vacController := NewVacancyController(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/parser/1", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/parser/{id}", vacController.Get)

	r.ServeHTTP(rr, req)

	var resp models.Vacancy
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(errors.ErrTestDecode)
	}

	assert.Equal(t, vacancy, resp)
}

func TestVacancyController_List_Internal(t *testing.T) {
	serviceMock := mocks.NewVacancyer(t)
	serviceMock.On("List").Return(nil, fmt.Errorf("test error"))

	vacController := NewVacancyController(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/parser/list", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/parser/list", vacController.List)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestVacancyController_List(t *testing.T) {
	expect := []models.Vacancy{vacancy}

	serviceMock := mocks.NewVacancyer(t)
	serviceMock.On("List").Return(expect, nil)

	vacController := NewVacancyController(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/parser/list", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/parser/list", vacController.List)

	r.ServeHTTP(rr, req)

	var resp []models.Vacancy

	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(errors.ErrTestDecode)
	}

	assert.Equal(t, expect, resp)
}
