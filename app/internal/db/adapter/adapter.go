package adapter

import (
	"encoding/json"
	"fmt"
	"log"
	"projects/LDmitryLD/parser/app/internal/infrastructure/errors"
	"projects/LDmitryLD/parser/app/internal/models"

	"github.com/jmoiron/sqlx"
)

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name=SQLAdapterer
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
}

func (s *SQLAdapter) Insert(vac models.Vacancy) (int, error) {
	var jobLocRaw []byte
	var err error

	switch jobLoc := vac.JobLocation.(type) {
	case map[string]interface{}:
		jobLocRaw, err = json.Marshal(jobLoc)
		if err != nil {
			log.Println("ошибка при маршалинге:", err)
			return 0, fmt.Errorf("ошибка при маршалинге: %s", err.Error())
		}
	case []interface{}:
		jobLocRaw, err = json.Marshal(jobLoc)
		if err != nil {
			log.Println("ошибка при маршалинге:", err)
			return 0, fmt.Errorf("ошибка при маршалинге: %s", err.Error())
		}
	default:
		log.Println("не удалось записать поле JobLocation в кэш")
	

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

	return id, nil
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
