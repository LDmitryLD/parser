package responder

import (
	"encoding/json"
	"log"
	"net/http"
	"projects/LDmitryLD/parser/app/internal/models"
)

type Responder interface {
	OutputJSON(w http.ResponseWriter, responseData interface{})

	ErrInternal(w http.ResponseWriter, err error)
	ErrBadRequest(w http.ResponseWriter, err error)
	ErrNotFound(w http.ResponseWriter, err error)
}

type Respond struct{}

func NewResponder() Responder {
	return &Respond{}
}

func (r *Respond) OutputJSON(w http.ResponseWriter, responseData interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		log.Println("ошибка при отправке ответа:", err)
		r.ErrInternal(w, err)
	}
}

func (r *Respond) ErrInternal(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	message := models.ApiResponse{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	}

	w.WriteHeader(http.StatusInternalServerError)

	if err := json.NewEncoder(w).Encode(message); err != nil {
		log.Println(err)
	}
}

func (r *Respond) ErrBadRequest(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	message := models.ApiResponse{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}

	w.WriteHeader(http.StatusBadRequest)

	if err := json.NewEncoder(w).Encode(message); err != nil {
		log.Println(err)
	}
}

func (r *Respond) ErrNotFound(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	message := models.ApiResponse{
		Code:    http.StatusNotFound,
		Message: err.Error(),
	}

	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		log.Println(err)
	}
}
