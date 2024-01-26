package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"projects/LDmitryLD/parser/app/internal/infrastructure/errors"
	"projects/LDmitryLD/parser/app/internal/infrastructure/responder"
	"projects/LDmitryLD/parser/app/internal/models"
	"projects/LDmitryLD/parser/app/internal/modules/vacancy/service"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Vacancyer interface {
	Search(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type VacancyController struct {
	service service.Vacancyer
	responder.Responder
}

func NewVacancyController(service service.Vacancyer) Vacancyer {
	return &VacancyController{
		service:   service,
		Responder: &responder.Respond{},
	}
}

func (v *VacancyController) Search(w http.ResponseWriter, r *http.Request) {
	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(LogErrDecodeReq)
		v.ErrBadRequest(w, err)
		return
	}

	vacs, err := v.service.Search(req.Query)
	if err != nil {
		v.ErrInternal(w, err)
		return
	}

	v.OutputJSON(w, vacs)
}

func (v *VacancyController) Delete(w http.ResponseWriter, r *http.Request) {
	idRaw := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		v.ErrBadRequest(w, err)
		return
	}

	if err := v.service.Delete(id); err != nil {
		if err != errors.ErrNotFound {
			v.ErrInternal(w, err)
			return
		} else {
			v.ErrNotFound(w, err)
			return
		}
	}

	resp := models.ApiResponse{
		Code:    200,
		Message: fmt.Sprintf("вакансия с id %d удалена", id),
	}

	v.OutputJSON(w, resp)
}

func (v *VacancyController) Get(w http.ResponseWriter, r *http.Request) {

	idRaw := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		v.ErrBadRequest(w, err)
		return
	}

	vac, err := v.service.Get(id)
	if err != nil {
		switch {
		case err != errors.ErrNotFound:
			v.ErrInternal(w, err)
			return
		default:
			v.ErrNotFound(w, err)
			return
		}
	}

	v.OutputJSON(w, vac)
}

func (v *VacancyController) List(w http.ResponseWriter, r *http.Request) {
	vacs, err := v.service.List()
	if err != nil {
		v.ErrInternal(w, err)
		return
	}

	v.OutputJSON(w, vacs)
}
