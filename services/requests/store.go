package requests

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

func (s *Store) GetAllRequests() ([]types.Request, error) {
    ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
    defer cancel()

    query := `
        SELECT * FROM requests;
    `

    rows, err := s.db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }

    var requests []types.Request
    for rows.Next() {
        var request types.Request
        err := rows.Scan(
            &request.ID,
            &request.Status,
        )
        if err != nil {
            return nil, err
        }
        requests = append(requests, request)
    }
    return requests, nil
}

func (s *Store) GetRequestById(id string) (*types.Request, error) {
    ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
    defer cancel()

    query := `SELECT * from requests WHERE id = $1`

    var request types.Request
    row := s.db.QueryRowContext( ctx, query, id)
    err := row.Scan(
        &request.ID,
        &request.Status,
    )
    if err != nil {
        return nil, err
    }

    return &request, nil
}

func (s *Store) CreateRequest(request types.Request) error {
    ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
    defer cancel()

    query := `
        INSERT INTO requests(
            id,
            status
        )
        VALUES($1, $2);
    `

    _, err := s.db.ExecContext(
        ctx,
        query,
        request.ID,
        request.Status,
    )
    if err != nil {
        return err
    }

    return nil
}

func (s *Store) UpdateRequest(request types.Request) error {
    return nil
}
