package models

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
