package docs

import (
	"projects/LDmitryLD/parser/app/internal/models"
	"projects/LDmitryLD/parser/app/internal/modules/vacancy/controller"
)

// swagger:route POST /parser/search parser SearchRequest
// Поиск вакансий.
// responses:
// 	200: SearchResponse

// swagger:parameters SearchRequest
type SearchRequest struct {
	// Название вакансии
	//
	// requierd:true
	// in:body
	Body controller.SearchRequest `json:"query"`
}

// swagger:response SearchResponse
type SearchResponse struct {
	// in:body
	Body []models.Vacancy
}

// swagger:route DELETE /parser/delete/{id} parser DeleteRequest
// Удаление вакансии по ID.
// responses:
//  200: DeleteResponse

// swagger:parameters DeleteRequest
type DeleteRequest struct {
	// ID вакансии
	//
	// required:true
	// in:path
	ID int `json:"id"`
}

// swagger:response DeleteResponse
type DeleteResponse struct {
	// in:body
	Body models.ApiResponse `json:"body"`
}

// swagger:route GET /parser/{id} parser GetRequest
// Получение вакансии по ID.
// responses:
//  200: GetResponse

// swagger:parameters GetRequest
type GetRequest struct {
	// ID вакансии
	//
	// requierd:true
	// in:path
	ID string `json:"id"`
}

// swagger:response GetResponse
type GetResponse struct {
	// in:body
	Body models.Vacancy `json:"body"`
}

// swagger:route GET /parser/list parser ListRequest
// Получить список ваканский из имеющихся в базе.
// responses:
//  200: ListResponse

// swagger:response
type ListResponse struct {
	// in:body
	List []models.Vacancy `json:"list"`
}
