package subjects

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/arturfil/meetings_app_server/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAllSubjects() ([]types.SubjectResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	query := `
        SELECT s.id, s.name, c.name, s.description, s.created_at, s.updated_at 
        FROM subjects s
        INNER JOIN subject_categories c ON c.id = s.category_id;
    `

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var subjects []types.SubjectResponse

	for rows.Next() {
		var subject types.SubjectResponse
		err := rows.Scan(
			&subject.ID,
			&subject.Name,
			&subject.Category,
			&subject.Description,
			&subject.CreatedAt,
			&subject.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		subjects = append(subjects, subject)
	}

	return subjects, nil
}

func (s *Store) SearchSubject(queryWord string) ([]types.SubjectResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	query := `
    SELECT s.id, s.name, c.name, s.description, s.created_at, s.updated_at FROM subjects s
        JOIN subject_categories c ON c.id = s.category_id
        WHERE (
            LOWER(s.name) LIKE LOWER($1) OR
            LOWER(c.name) LIKE LOWER($1) OR
            LOWER(s.description) LIKE LOWER($1)
        );
    `

    searchPattern := fmt.Sprintf("%%%s%%", queryWord)

	rows, err := s.db.QueryContext(ctx, query, searchPattern)
	if err != nil {
	    return nil, err
	}

    var subjects []types.SubjectResponse

    for rows.Next() {
        var subject types.SubjectResponse
        err := rows.Scan(
            &subject.ID,
            &subject.Name,
            &subject.Category,
            &subject.Description,
            &subject.CreatedAt,
            &subject.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        subjects = append(subjects, subject)
    }

	return subjects, nil
}

func (s *Store) GetAllSubjectsByCategory(categoryId string) ([]types.SubjectResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	query := `
        select s.id, s.name, c.name, s.description, s.created_at, s.updated_at from subjects s
        LEFT JOIN subject_categories c ON s.category_id = c.id
        where s.category_id = $1;
    `

	rows, err := s.db.QueryContext(ctx, query, categoryId)
	if err != nil {
		return nil, err
	}

	var subjects []types.SubjectResponse

	for rows.Next() {
		var subject types.SubjectResponse
		err := rows.Scan(
			&subject.ID,
			&subject.Name,
			&subject.Category,
			&subject.Description,
			&subject.CreatedAt,
			&subject.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		subjects = append(subjects, subject)
	}

	return subjects, nil
}

func (s *Store) GetSubjectById(id string) (types.Subject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	query := `
        select * from subjects
        where id = $1;
    `

	var subject types.Subject
	row := s.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&subject.ID,
		&subject.Name,
		&subject.Description,
		&subject.CategoryID,
		&subject.CreatedAt,
		&subject.UpdatedAt,
	)
	if err != nil {
		return types.Subject{}, err
	}

	return subject, nil
}

func (s *Store) GetSubjectCategories() ([]types.SubjectCategory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	query := `
        SELECT id, name, description, created_at, updated_at
        FROM subject_categories
        ORDER BY name
    `

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []types.SubjectCategory

	for rows.Next() {
		var category types.SubjectCategory
		if err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
			&category.UpdatedAt,
		); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
