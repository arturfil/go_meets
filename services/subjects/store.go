package subjects

import (
	"context"
	"database/sql"

	"github.com/arturfil/meetings_app_server/types"
)

type Store struct {
    db *sql.DB
}

func NewStore(db *sql.DB) *Store {
    return &Store{db: db}
}

func (s *Store) GetAllSubjects() ([]types.Subject, error) {
    ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
    defer cancel()

    query := `
        SELECT * FROM subjects;
    `

    rows, err := s.db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }

    var subjects []types.Subject

    for rows.Next() {
        var subject types.Subject
        err := rows.Scan(
            &subject.ID,
            &subject.Name,
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

