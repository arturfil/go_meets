package types

import "time"

type TeachingsStore interface {
	GetAllTeachings(userId string) ([]Teaching, error)
	CreateTeaching(teaching TeachingSubmission) error
	GetSchedules(userId string) ([]Schedule, error)
	CreateSchedule(schedule Schedule) error
	DeleteSchedule(userId string) error
}

type Teaching struct {
	ID        string    `json:"id"`
	TeacherId string    `json:"teacher_id"`
	SubjectId string    `json:"subject_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TeachingSubmission struct {
    TeacherID string `json:"teacher_id"`
    SubjectID string `json:"subject_id"`
}

type Schedule struct {
	UserId    string    `json:"user_id"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
	Day       string    `json:"day"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Testing struct {
	Added     bool
	Done      bool
	Tests     int
	CreatedAt time.Time
}
