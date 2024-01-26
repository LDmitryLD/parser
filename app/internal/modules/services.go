package modules

import (
	"projects/LDmitryLD/parser/app/internal/modules/vacancy/service"
	"projects/LDmitryLD/parser/app/internal/parser"
	"projects/LDmitryLD/parser/app/internal/storages"
)

type Services struct {
	Vacancy service.Vacancyer
}

func NewServices(storages *storages.Storages, parser *parser.SeleniumParser) *Services {
	vacService := service.NewVacanceService(parser, storages.Vacancy)

	return &Services{
		Vacancy: vacService,
	}
}
