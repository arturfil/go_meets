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
    row, err := s.db.QueryContext( ctx, query, email)
    if err != nil {
        return nil, err
    }

    err = row.Scan(
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
    return &types.User{}, nil
}

func (s *Store) SignUp(user types.User) (*types.User, error) {
    return &types.User{}, nil
}

func (s *Store) PasswordMatches(plainText string) (bool, error) {
    return false, nil
}

func (s *Store) Update() error {
    return nil
}

func (s *Store) Delete(id string) error {
    return nil
}
