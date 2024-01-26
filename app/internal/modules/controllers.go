package modules

import "projects/LDmitryLD/parser/app/internal/modules/vacancy/controller"

type Controllers struct {
	Vacancy controller.Vacancyer
}

func NewControllers(services *Services) *Controllers {
	vaController := controller.NewVacancyController(services.Vacancy)

	return &Controllers{
		Vacancy: vaController,
	}
}
