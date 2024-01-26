package models

// type Vacancy struct {
// 	Context            string             `json:"@context" db:"context"`
// 	Type               string             `json:"@type" db:"type"`
// 	DatePosted         string             `json:"datePosted" db:"date_posted"`
// 	Title              string             `json:"title" db:"title"`
// 	Description        string             `json:"description" db:"description"`
// 	Identifier         Identifier         `json:"identifier" db:"identifier"`
// 	ValidThrough       string             `json:"validThrough" db:"valid_through"`
// 	HiringOrganization HiringOrganization `json:"hiringOrganization" db:"hiring_organization"`
// 	JobLocation        JobLocation        `json:"jobLocation" db:"job_location"`
// 	JobLocationType    string             `json:"jobLocationType" db:"job_location_type"`
// 	EmploymentType     string             `json:"employmentType" db:"employment_type"`
// }

// type HiringOrganization struct {
// 	Type   string `json:"@type" db:"type"`
// 	Name   string `json:"name" db:"name"`
// 	Logo   string `json:"logo" db:"logo"`
// 	SameAs string `json:"sameAs" db:"same_as"`
// }

// type Identifier struct {
// 	Type  string `json:"@type" db:"type"`
// 	Name  string `json:"name" db:"name"`
// 	Value string `json:"value"  db:"value"`
// }

// type JobLocation struct {
// 	Type    string  `json:"@type" db:"type"`
// 	Address Address `json:"address" db:"address"`
// }

// type Address struct {
// 	Type            string         `json:"@type" db:"type"`
// 	StreetAddress   string         `json:"streetAddress" db:"street_address"`
// 	AddressLocality string         `json:"addressLocality" db:"address_locality"`
// 	AddressCountry  AddressCountry `json:"addressCountry" db:"address_country"`
// }

// type AddressCountry struct {
// 	Type string `json:"@type" db:"type"`
// 	Name string `json:"name" db:"name"`
// }

type Vacancy struct {
	Context            string             `json:"@context,omitempty" db:"context"`
	Type               string             `json:"@type,omitempty" db:"type"`
	DatePosted         string             `json:"datePosted,omitempty" db:"date_posted"`
	Title              string             `json:"title,omitempty" db:"title"`
	Description        string             `json:"description,omitempty" db:"description"`
	Identifier         Identifier         `json:"identifier,omitempty" db:"identifier"`
	ValidThrough       string             `json:"validThrough,omitempty" db:"valid_through"`
	HiringOrganization HiringOrganization `json:"hiringOrganization,omitempty" db:"hiring_organization"`
	JobLocation        interface{}        `json:"jobLocation,omitempty" db:"job_location"`
	JobLocationType    string             `json:"jobLocationType" db:"job_location_type"`
	EmploymentType     string             `json:"employmentType,omitempty" db:"employment_type"`
}

type HiringOrganization struct {
	Type   string `json:"@type" db:"type"`
	Name   string `json:"name" db:"name"`
	Logo   string `json:"logo" db:"logo"`
	SameAs string `json:"sameAs" db:"same_as"`
}

type Identifier struct {
	Type  string `json:"@type" db:"type"`
	Name  string `json:"name" db:"name"`
	Value string `json:"value" db:"value"`
}

type JobLocation struct {
	Type    string `json:"@type"`
	Address string `json:"address"`
}

type PropertyValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Organization struct {
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	SameAs string `json:"sameAs"`
}

// type PostalAddress struct {
// 	StreetAddress   string `json:"streetAddress"`
// 	AddressLocality string `json:"addressLocality"`
// 	AddressCountry  struct {
// 		Name string `json:"name"`
// 	} `json:"addressCountry"`
// }

// type Place struct {
// 	Address PostalAddress `json:"address"`
// }

// type Vacancy struct {
// 	Context            string        `json:"@context"`
// 	Type               string        `json:"@type"`
// 	DatePosted         string        `json:"datePosted"`
// 	Title              string        `json:"title"`
// 	Description        string        `json:"description"`
// 	Identifier         PropertyValue `json:"identifier"`
// 	ValidThrough       string        `json:"validThrough"`
// 	HiringOrganization Organization  `json:"hiringOrganization"`
// 	JobLocation        Place         `json:"jobLocation"`
// 	JobLocationType    string        `json:"jobLocationType"`
// 	EmploymentType     string        `json:"employmentType"`
// }

// type Vacancy struct {
// 	Context            string       `json:"@context"`
// 	Type               string       `json:"@type"`
// 	DatePosted         string       `json:"datePosted"`
// 	Title              string       `json:"title"`
// 	Description        string       `json:"description"`
// 	Identifier         Identifier   `json:"identifier"`
// 	ValidThrough       string       `json:"validThrough"`
// 	HiringOrganization Organization `json:"hiringOrganization"`
// 	JobLocation        []Place      `json:"jobLocation"`
// 	JobLocationType    string       `json:"jobLocationType"`
// 	EmploymentType     string       `json:"employmentType"`
// }

// type Identifier struct {
// 	Type  string `json:"@type"`
// 	Name  string `json:"name"`
// 	Value string `json:"value"`
// }

// type Organization struct {
// 	Type    string         `json:"@type"`
// 	Name    string         `json:"name"`
// 	Logo    string         `json:"logo"`
// 	SameAs  string         `json:"sameAs"`
// 	Address *PostalAddress `json:"address"`
// }

// type Place struct {
// 	Type    string      `json:"@type"`
// 	Address interface{} `json:"address"` // Используем интерфейс, так как поле может быть строкой или структурой
// }

//

//

//

// type PostalAddress struct {
// 	Type            string   `json:"@type"`
// 	StreetAddress   string   `json:"streetAddress"`
// 	AddressLocality string   `json:"addressLocality"`
// 	AddressCountry  *Country `json:"addressCountry"`
// }

// type Country struct {
// 	Type string `json:"@type"`
// 	Name string `json:"name"`
// }

// type Vacancy struct {
// 	Context            string       `json:"@context"`
// 	Type               string       `json:"@type"`
// 	DatePosted         string       `json:"datePosted"`
// 	Title              string       `json:"title"`
// 	Description        string       `json:"description"`
// 	Identifier         Identifier   `json:"identifier"`
// 	ValidThrough       string       `json:"validThrough"`
// 	HiringOrganization Organization `json:"hiringOrganization"`
// 	JobLocation        []Place      `json:"jobLocation"`
// 	JobLocationType    string       `json:"jobLocationType"`
// 	EmploymentType     string       `json:"employmentType"`
// }

// type VacancyJSON struct {
// 	Vacancies []struct {
// 		Vacancy json.RawMessage `json:"vacancy"`
// 	} `json:"vacancies"`
// }

// type Place struct {
// 	Type    string          `json:"@type"`
// 	Address json.RawMessage `json:"address"`
// }
