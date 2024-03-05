package services

import "time"

type User struct {
    ID string `json:"id"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
	CreatedAt   time.Time `json:"createdat"`
	UpdatedAt   time.Time `json:"updatedat"`
}
