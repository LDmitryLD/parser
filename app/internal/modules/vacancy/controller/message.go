package controller

const (
	LogErrDecodeReq = "ошибка при декодировании запроса"
)

type SearchRequest struct {
	Query string `json:"query"`
}
