package adapter

import (
	"encoding/json"
	"fmt"
	"log"
	"projects/LDmitryLD/parser/app/internal/infrastructure/errors"
	"projects/LDmitryLD/parser/app/internal/models"
	"reflect"

	"github.com/jmoiron/sqlx"
)

type SQLAdapterer interface {
	SelectVacancies(query string, list bool) ([]models.Vacancy, error)
	SelectVacancy(id int) (models.Vacancy, error)
	Insert(vac models.Vacancy) (int, error)
	Delete(id int) error
}

type SQLAdapter struct {
	db *sqlx.DB
}

func NewSQLAdapter(db *sqlx.DB) *SQLAdapter {
	return &SQLAdapter{
		db: db,
	}
}

func (s *SQLAdapter) selectQuery(query string) ([]models.Vacancy, error) {
	q := `
	SELECT 
		v.context, v.type, v.date_posted, v.title, v.description, v.valid_through, v.job_location,v.job_location_type, v.employment_type,
		h.type, h.name, h.logo, h.same_as, 
		i.type, i.name, i.value
	FROM
		vacancy v
		JOIN hiring_organization h ON v.id = h.vacancy_id
		JOIN identifier i ON v.id = i.vacancy_id
	WHERE 
		description ILIKE '%' || $1 || '%'
	`

	rows, err := s.db.Query(q, query)
	if err != nil {
		log.Println("ошибка при получении вакансии из бд:", err)
		return nil, err
	}

	if !rows.Next() {
		return nil, errors.ErrNotFound
	}

	var vacs []models.Vacancy

	for rows.Next() {
		var v models.Vacancy
		err := rows.Scan(&v.Context, &v.Type, &v.DatePosted, &v.Title, &v.Description, &v.ValidThrough, &v.JobLocation, &v.JobLocationType, &v.EmploymentType,
			&v.HiringOrganization.Type, &v.HiringOrganization.Name, &v.HiringOrganization.Logo, &v.HiringOrganization.SameAs,
			&v.Identifier.Type, &v.Identifier.Name, &v.Identifier.Value,
		)
		if err != nil {
			log.Println("ошибка при сканировании результата:", err)
			return nil, err
		}
		vacs = append(vacs, v)
	}

	return vacs, nil
}

func (s *SQLAdapter) SelectList() ([]models.Vacancy, error) {
	q := `
	SELECT
		v.context, v.type, v.date_posted, v.title, v.description, v.valid_through, v.job_location, v.job_location_type, v.employment_type,
		h.type, h.name, h.logo, h.same_as,
		i.type, i.name, i.value
	FROM
		vacancy v
		JOIN hiring_organization h ON v.id = h.vacancy_id
		JOIN identifier i ON v.id = i.vacancy_id	
	`

	rows, err := s.db.Query(q)
	if err != nil {
		log.Println("ошибка при получении всех вакансии из бд:", err)
		return nil, err
	}

	if !rows.Next() {
		return nil, errors.ErrNotFound
	}

	var vacs []models.Vacancy

	for rows.Next() {
		var v models.Vacancy
		err := rows.Scan(&v.Context, &v.Type, &v.DatePosted, &v.Title, &v.Description, &v.ValidThrough, &v.JobLocation, &v.JobLocationType, &v.EmploymentType,
			&v.HiringOrganization.Type, &v.HiringOrganization.Name, &v.HiringOrganization.Logo, &v.HiringOrganization.SameAs,
			&v.Identifier.Type, &v.Identifier.Name, &v.Identifier.Value,
		)
		if err != nil {
			log.Println("ошибка при сканировании результата:", err)
			return nil, err
		}
		vacs = append(vacs, v)
	}

	return vacs, nil
}

func (s *SQLAdapter) SelectVacancies(query string, list bool) ([]models.Vacancy, error) {

	if !list {
		return s.selectQuery(query)
	} else {
		return s.SelectList()
	}

}

// SELECT
// 		v.context, v.type, v.date_posted, v.title, v.description, v.valid_through, v.job_location_type, v.employment_type,
// 		h.type, h.name, h.logo, h.same_as,
// 		i.type, i.name, i.value,
// 		jl.type,
// 		a.type, a.street_address, a.address_locality,
// 		ac.type, ac.name
// 	FROM
// 		vacancy v
// 		JOIN hiring_organization h ON v.id = h.vacancy_id
// 		JOIN identifier i ON v.id = i.vacancy_id
// 		JOIN job_location jl ON v.id = jl.vacancy_id
// 		JOIN address a ON jl.id = a.job_location_id
// 		JOIN address_country ac ON a.id = ac.address_id
// 	`

// rows, err := s.db.Query(q, query)
// if err != nil {
// 	log.Println("ошибка при запросе в бд списка вакансий:", err)
// 	return nil, err
// }

// if !rows.Next() {
// 	log.Println("совпадений нет")
// 	return nil, fmt.Errorf("совпадений нет")
// }

// var vacs []models.Vacancy

// for rows.Next() {
// 	var vac models.Vacancy
// 	err := rows.Scan(&vac.Context, &vac.Type, &vac

// 	)
// }

func (s *SQLAdapter) SelectVacancy(id int) (models.Vacancy, error) {

	q := `
		SELECT
			type, name, logo, same_as
		FROM
			hiring_organization h
		WHERE
			h.id = $1
		LIMIT 1		
	`

	var hirOrg models.HiringOrganization

	if err := s.db.Get(&hirOrg, q, id); err != nil {
		log.Println("ошибка при получении hiring_organization из дб:", err)
	}

	q = `
		SELECT 
			type, name, value
		FROM
			identifier i
		WHERE 
			i.id = $1
		LIMIT 1			
	`

	var ident models.Identifier

	if err := s.db.Get(&ident, q, id); err != nil {
		log.Println("ошибка при получении  identifier из бд", err)
	}

	q = `
		SELECT 
			context, type, date_posted, title, description, valid_through, job_location, job_location_type, employment_type
		FROM
			vacancy v
		WHERE
			v.id = $1
		LIMIT 1	
	`
	var vac models.Vacancy

	if err := s.db.Get(&vac, q, id); err != nil {
		log.Println("ошибка s.db.Get()", err)
		return models.Vacancy{}, err
	}

	vac.HiringOrganization = hirOrg
	vac.Identifier = ident

	return vac, nil

	// q := `
	// 	SELECT
	// 		v.context, v.type, v.date_posted, v.title, v.description, v.valid_through, v.job_location, v.job_location_type, v.employment_type,
	// 		h.type, h.name, h.logo, h.same_as,
	// 		i.type, i.name, i.value
	// 	FROM
	// 		vacancy v
	// 		JOIN hiring_organization h ON v.id = h.vacancy_id
	// 		JOIN identifier i ON v.id = i.vacancy_id
	// 	WHERE
	// 		v.id = $1
	// 	LIMIT 1
	// `
	// var vac models.Vacancy

	// if err := s.db.Get(&vac, q, id); err != nil {
	// 	log.Println("ошибка s.db.Get()", err)
	// 	return models.Vacancy{}, err
	// }

	// return vac, nil
}

func (s *SQLAdapter) Insert(vac models.Vacancy) (int, error) {
	var jobLocRaw []byte
	var err error
	// попробовать переделать с if на type switch
	// ну или вообще сначала попробовать маршалить не приведённое значение)
	jobLocMap, ok := vac.JobLocation.(map[string]interface{})
	if !ok {
		log.Println("не удалось преобразовать, cейчас будт пробовать в слайс, настоящее значение типа:", reflect.TypeOf(vac.JobLocation))
		jobLocArr, ok := vac.JobLocation.([]interface{})
		if !ok {
			log.Println("в слайс тоже не получилось")
		}

		jobLocRaw, err = json.Marshal(jobLocArr)
		if err != nil {
			log.Println("ошибка при маршалинге []interface{}, ", err)
		}

	} else {
		jobLocRaw, err = json.Marshal(jobLocMap)
		if err != nil {
			log.Println("Ошибка при маршалинге:", err)
			return 0, fmt.Errorf("ошибка при маршалинге")
		}
	}

	q := `
		INSERT INTO vacancy
			(context, type, date_posted, title, description, valid_through, job_location, job_location_type, employment_type)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`
	var id int
	if err := s.db.QueryRow(q, vac.Context, vac.Type, vac.DatePosted, vac.Title, vac.Description, vac.ValidThrough, string(jobLocRaw), vac.JobLocationType, vac.EmploymentType).Scan(&id); err != nil {
		log.Println("ошибка при записи в таблицу: ", err)
		return 0, err
	}

	hiringOrg := vac.HiringOrganization

	q = `
		INSERT INTO hiring_organization
			(vacancy_id, type, name, logo, same_as)
		VALUES
			($1, $2, $3, $4, $5)
	`
	_, err = s.db.Exec(q, id, hiringOrg.Type, hiringOrg.Name, hiringOrg.Logo, hiringOrg.SameAs)
	if err != nil {
		log.Println("ошибка при записи в таблицу: ", err)
		return 0, err
	}

	ident := vac.Identifier

	q = `
		INSERT INTO identifier
			(vacancy_id, type, name, value)
		VALUES
			($1, $2, $3, $4)
	`

	_, err = s.db.Exec(q, id, ident.Type, ident.Name, ident.Value)
	if err != nil {
		log.Println("ошибка при записи в таблицу: ", err)
		return 0, err
	}
	log.Println("данные записаны в бд! ID: ", id)
	return id, nil

	//	jobLoc := vac.JobLocation

	// q = `
	// 	INSERT INTO job_location
	// 		(vacancy_id, type)
	// 	VALUES
	// 		($1, $2)
	// `

	// _, err = s.db.Exec(q, id, jobLoc.Type)
	// if err != nil {
	// 	log.Println("ошибка при записи в таблицу: ", err)
	// 	return 0, err
	// }

	//addr := vac.JobLocation.Address

	// q = `
	// 	INSERT INTO address
	// 		(job_location_id, type, street_address, address_locality)
	// 	VALUES
	// 		($1, $2, $3, $4)
	// `

	// _, err = s.db.Exec(q, id, addr.Type, addr.StreetAddress, addr.AddressLocality)
	// if err != nil {
	// 	log.Println("ошибка при записи в таблицу: ", err)
	// 	return 0, err
	// }

	// addrCountry := addr.AddressCountry

	// q = `
	// 	INSERT INTO address_country
	// 		(address_id, type, name)
	// 	VALUES
	// 		($1, $2, $3)
	// `

	// // _, err = s.db.Exec(q, id, addrCountry.Type, addrCountry.Name)
	// // if err != nil {
	// // 	log.Println("ошибка при записи в таблицу: ", err)
	// // 	return 0, err
	// // }

	// return id, nil

	//panic("emplement me ")
}

func (s *SQLAdapter) Delete(id int) error {
	q := `
		DELETE FROM vacancy
		WHERE id = ($1)
	`
	_, err := s.db.Exec(q, id)
	if err != nil {
		fmt.Println("ошибка при удалении:", err)
		return err
	}

	return nil
}
