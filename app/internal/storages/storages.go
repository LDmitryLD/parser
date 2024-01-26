package storages

import (
	"projects/LDmitryLD/parser/app/internal/db/adapter"
	"projects/LDmitryLD/parser/app/internal/modules/vacancy/storage"
)

type Storages struct {
	Vacancy storage.VacancyStorager
}

func NewStorages(sqlAdapter *adapter.SQLAdapter) *Storages {
	return &Storages{
		Vacancy: storage.NewVacancyStorage(sqlAdapter),
	}
}
