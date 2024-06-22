package types

import "github.com/golang-jwt/jwt/v4"

type UserStore interface {
	GetTeachers() ([]User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserById(id string) (*User, error)
	CreateUser(user RegisterUserPayload) error
	Update() error
	Delete(id string) error
}

type UsersAndRoles struct {
	UserID      string `json:"user_id"`
	Email       string `json:"email"`
	RoleID      string `json:"role_id"`
	Description string `json:"description"`
}

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"-"`
	IsAdmin   string `json:"is_admin"`
	IsTeacher string `json:"is_teacher"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Register Payload Types here
type RegisterUserPayload struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=100"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=100"`
}

type TokenClaim struct {
	Sub string `json:"sub"`
	Aud string `json:"aud"`
	Iss string `json:"iss"`
	Exp string `json:"exp"`
	jwt.RegisteredClaims
}
