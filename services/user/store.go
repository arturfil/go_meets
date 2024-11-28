package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/arturfil/meetings_app_server/types"
	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAllUsers() ([]types.UserReponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	query := `
        SELECT
            id, email, first_name, last_name, created_at, updated_at 
        FROM users 
    `

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var users []types.UserReponse
	for rows.Next() {
		var user types.UserReponse
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		roles, err := s.getUserRoles(user.ID)
		if err != nil {
			return nil, err
		}

		user.Roles = roles
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (s *Store) GetUserByEmail(email string) (*types.UserReponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	query := `
            SELECT
                id, email, first_name, last_name, password, created_at, updated_at 
            FROM users 
            WHERE email = $1
        `

	var user types.UserReponse

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

	roles, err := s.getUserRoles(user.ID)
	if err != nil {
		return nil, err
	}

	user.Roles = roles

	return &user, nil
}

func (s *Store) GetUserById(id string) (*types.UserReponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	query := `
        SELECT u.id, u.email, u.first_name, u.created_at, u.updated_at 
        FROM users u where id = $1
    `

	var user types.UserReponse

	row := s.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	roles, err := s.getUserRoles(user.ID)
	if err != nil {
		return nil, err
	}

	user.Roles = roles

	return &user, nil
}

func (s *Store) CreateUser(user types.RegisterUserPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	newId := uuid.New()

	query := `
        INSERT INTO users (id, first_name, last_name, email, password)
        VALUES ($1, $2, $3, $4, $5);
    `

	_, err := s.db.ExecContext(
		ctx,
		query,
		newId,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password, // already encrypted at this point
	)
	if err != nil {
		return err
	}

	role_query := `
        INSERT INTO role_relations (user_id, role_id)
        VALUES ($1, $2);
    `

	_, err = s.db.ExecContext(
		ctx,
		role_query,
		newId,
		"22b3f2ca-3e98-447f-a807-9609fa496ae9", // user role id,
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
        SELECT u.id, email, first_name, last_name, created_at, updated_at
        FROM users u 
        JOIN role_relations rr ON rr.user_id = u.id
        JOIN roles r ON r.id = rr.role_id
        WHERE r.description = 'teacher';
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

func (s *Store) SearchTeachers(searchQuery string) ([]types.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	query := `
        SELECT DISTINCT u.id, u.first_name, u.last_name, u.email, u.created_at, u.updated_at
        FROM users u
        INNER JOIN role_relations rr ON u.id = rr.user_id
        INNER JOIN roles r ON rr.role_id = r.id
        WHERE r.description = 'teacher'
        AND (
            LOWER(u.first_name) LIKE LOWER($1) OR
            LOWER(u.last_name) LIKE LOWER($1) OR
            LOWER(u.email) LIKE LOWER($1)
        )
    `

	searchPattern := fmt.Sprintf("%%%s%%", searchQuery)

	rows, err := s.db.QueryContext(ctx, query, searchPattern)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []types.User
	for rows.Next() {
		var teacher types.User
		err := rows.Scan(
			&teacher.ID,
			&teacher.FirstName,
			&teacher.LastName,
			&teacher.Email,
			&teacher.CreatedAt,
			&teacher.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return teachers, nil
}

func (s *Store) Update() error {
	return nil
}

func (s *Store) Delete(id string) error {
	return nil
}

func (s *Store) getUserRoles(userId string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	query := `
        SELECT r.description FROM users u
        LEFT JOIN role_relations rr ON u.id = rr.user_id
        LEFT JOIN roles r ON r.id = rr.role_id 
        WHERE u.id = $1
    `

	rows, err := s.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	var roles []string
	for rows.Next() {
		var role string
		err := rows.Scan(&role)

		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}
