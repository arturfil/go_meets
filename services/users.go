package services

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsAdmin   bool      `json:"is_admin"`
	IsTeacher string    `json:"is_teacher"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

// GetByEmail - this method returns user by providing an email
func (u *User) GetByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        SELECT
            id, first_name, last_name, email, password, is_teacher, is_admin, created_at, updated_at
        FROM users
        WHERE email = $1
    `

	var user User

	row := db.QueryRowContext(ctx, query, email)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
        &user.IsAdmin,
		&user.IsTeacher,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (u *User) Signup(user User) (*User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
    defer cancel()

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
    if err != nil {
        return nil, err
    }

    query := `
        INSERT INTO users (first_name, last_name, email, password, is_admin, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7) returning *;
    `

    _, err = db.ExecContext(
        ctx, 
        query,
        user.FirstName,
        user.LastName,
        user.Email,
        hashedPassword,
        user.IsAdmin,
        time.Now(),
        time.Now(),
    )
    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (u *User) PasswordMatches(plainText string) (bool, error) {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
    if err != nil {
        switch {
            case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
                return false, nil
        default:
            return false, err
        }
    }
    return true, nil
}
