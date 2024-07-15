package types

import "time"

type TeachingsStore interface {
	GetAllTeachings(userId string) ([]Teaching, error)
	CreateTeaching(userId, subjectId string) error
	GetSchedule(userId string) (*Schedule, error)
}

type Teaching struct {
	ID        string    `json:"id"`
	TeacherId string    `json:"teacher_id"`
	SubjectId string    `json:"subject_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Schedule struct {
	UserId    string    `json:"user_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
    DayOfWeek time.Time `json:"day"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Testing struct {
	Added     bool
	Done      bool
	Tests     int
	CreatedAt time.Time
}
