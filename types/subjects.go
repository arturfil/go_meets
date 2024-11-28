package types

import "time"

type SubjectStore interface {
	GetAllSubjects() ([]SubjectResponse, error)
	GetSubjectCategories() ([]SubjectCategory, error)
	SearchSubject(queryWord string) ([]SubjectResponse, error)
	GetAllSubjectsByCategory(categoryId string) ([]SubjectResponse, error)
	GetSubjectById(id string) (Subject, error)
}

type Subject struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryID  string    `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SubjectResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SubjectCategory struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
