package roles

import (
	"database/sql"
	"fmt"
)

type Store struct {
    db *sql.DB
}

func NewStore(db *sql.DB) *Store {
    return &Store{db: db}
}

func (s *Store) GetRolesFromUser(id string) error {
    query := `
        SELECT u.id, u.first_name, u.last_name, r.description FROM users u
        LEFT JOIN role_relations rr ON rr.user_id = u.id
        LEFT JOIN roles r ON rr.role_id = r.id 
        WHERE u.id = $1;
    `

    fmt.Println("query:", query)

    return nil
}
