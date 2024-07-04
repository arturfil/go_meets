package requests

import (
	"context"
	"database/sql"
	"errors"

	"github.com/arturfil/meetings_app_server/types"
)

type Store struct {
    db *sql.DB
}

func NewStore(db *sql.DB) *Store {
    return &Store{db: db}
}

func (s *Store) GetAllRequests() ([]types.RequestResponse, error) {
    ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
    defer cancel()

    query := `
        SELECT r.id, u.first_name, u.last_name, r.status, r.type FROM requests r
        JOIN users u ON u.id = r.id;
    ;
    `

    rows, err := s.db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }

    var requests []types.RequestResponse
    for rows.Next() {
        var request types.RequestResponse
        err := rows.Scan(
            &request.ID,
            &request.FirstName,
            &request.LastName,
            &request.Status,
            &request.Type,
        )
        if err != nil {
            return nil, err
        }
        requests = append(requests, request)
    }
    return requests, nil
}

func (s *Store) GetRequestById(id string) (*types.RequestResponse, error) {
    ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
    defer cancel()

    query := `
        SELECT r.id, r.status, u.first_name, u.last_name, r.type, r.subject_request_type, r.subject_request_name
        FROM requests r 
        INNER JOIN users u ON u.id = r.id
        WHERE r.id = $1
    `

    var request types.RequestResponse
    row := s.db.QueryRowContext( ctx, query, id)
    err := row.Scan(
        &request.ID,
        &request.Status,
        &request.FirstName,
        &request.LastName,
        &request.Type,
        &request.SubjectRequestType,
        &request.SubjectRequestName,
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

func (s *Store) UpdateRequest(userId string, request types.Request) error {
    ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
    defer cancel()
   
    err := s.hasRole(userId, "admin")
    if err != nil {
        return err
    }

    err = s.hasRole(request.ID, "teacher")
    if err == nil {
        return errors.New("User already has this priviledge")
    }

    role_relation_query := `
        INSERT INTO role_relations
        (user_id, role_id)
        VALUES ($1, $2)
    `

    if request.Status == "approved" && request.Type == "teach request" {
        _, err = s.db.ExecContext(
            ctx,
            role_relation_query,
            &request.ID,
            "71dc50c1-1934-4da1-91a5-2fb73fadb39e", // fixed id of teacher role
        )
        if err != nil {
            return err
        }
    }

    query := `
        UPDATE requests
        SET status = $1
        WHERE id = $2
    `

    _, err = s.db.ExecContext(
        ctx,
        query,
        &request.Status, 
        &request.ID,
    )
    if err != nil {
        return err
    }

    return nil
}

func (s *Store) hasRole(id, roleAuth string) error {
    ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
    defer cancel()

    query := `
        SELECT u.id, u.email, r.id, r.description FROM users u 
        LEFT JOIN role_relations rr ON rr.user_id = u.id
        JOIN roles r ON rr.role_id = r.id 
        WHERE u.id = $1;
    `
    
    rows, err := s.db.QueryContext(ctx, query, id)
    if err != nil {
        return err 
    }

    var user_relations []types.UsersAndRoles
    for rows.Next() {
        var user types.UsersAndRoles

        err := rows.Scan(
            &user.UserID,
            &user.RoleID, 
            &user.Email, 
            &user.Description, 
        )
        if err != nil {
            return err
        }
        user_relations = append(user_relations, user)
    }

    var hasRole = false
    for _, role := range user_relations {
        if role.Description == roleAuth {
            hasRole = true
        }
    }

    if hasRole == false {
       return errors.New("user is not have the right permission to exectute his process")
    }

    return nil
}
