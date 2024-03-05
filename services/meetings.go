package services

import "time"

type Meeting struct {
	ID          string    `json:"id"`
	DateAndTime time.Time `json:"dateandtime"`
	Topic       string    `json:"topic"`
	StudentID   string    `json:"studentid"`
	TeacherID   string    `json:"teacherid"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"createdat"`
	UpdatedAt   time.Time `json:"updatedat"`
}
