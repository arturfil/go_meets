package user

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

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
    defer cancel()

    query := `
        SELECT
            id, email, first_name, last_name, password, created_at, updated_at 
		FROM users 
		WHERE email = $1
    `

    var user types.User
     
    // scan user
    row := s.db.QueryRowContext(ctx, query, email)
    err := row.Scan(
        &user.ID,
        &user.Email,
        &user.FirstName,
        &user.LastName,
        &user.Password,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (s *Store) GetUserById(id string) (*types.User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
    defer cancel()

    query := `select id, is_admin, email, first_name, password, created_at, updated_at from users where id = $1`

    var user types.User
    
    row := s.db.QueryRowContext( ctx, query, id)
    err := row.Scan(
        &user.ID,
        &user.IsAdmin,
        &user.Email,
        &user.FirstName,
        &user.Password,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (s *Store) CreateUser(user types.RegisterUserPayload) (error) {
    ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
    defer cancel()

    query := `
        INSERT INTO users (first_name, last_name, email, password)
        VALUES ($1, $2, $3, $4)
    ` 

    _, err := s.db.QueryContext(
        ctx,
        query,
        user.FirstName,
        user.LastName,
        user.Email,
        user.Password, // already encrypted at this point
    )
    if err != nil {
        return err
    }


    return nil
}

func (s *Store) GetTeachers() ([]types.User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
    defer cancel()

    query := `
        SELECT id, email, first_name, last_name, is_teacher, is_admin, created_at, updated_at
        FROM users u WHERE u.is_teacher = TRUE;
    `
    
    rows, err := s.db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }

    var teachers []types.User
    for rows.Next() {
        var teacher types.User
        err := rows.Scan(
            &teacher.ID,
            &teacher.Email,
            &teacher.FirstName,
            &teacher.LastName,
            &teacher.IsTeacher,
            &teacher.IsAdmin,
            &teacher.CreatedAt,
            &teacher.UpdatedAt,

        )
        if err != nil {
            return nil, err
        }

        teachers = append(teachers, teacher)
    }

    return teachers, nil
}

func (s *Store) Update() error {
    return nil
}

func (s *Store) Delete(id string) error {
    return nil
}
