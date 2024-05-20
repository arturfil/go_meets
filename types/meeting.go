package types

import "time"

type MeetingStore interface {
	GetAllMeetings() ([]MeetingResponse, error)
	GetMeetingByID(id string) (*Meeting, error)
	CreateMeeting(meeting Meeting) error
	UpdateMeeting(meeting Meeting) error
	DeleteMeeting(id string) error
}

type Meeting struct {
	ID              string    `json:"id"`
	SubjectID       string    `json:"subject_id"`
	StudentID       string    `json:"student_id"`
	TeacherID       string    `json:"teacher_id"`
	StudentAttended bool      `json:"student_attended"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Register Payload types here
type MeetingResponse struct {
	ID              string    `json:"id"`
	Subject         string    `json:"subject,omitempty"`
	Student         string    `json:"student,omitempty"`
	Teacher         string    `json:"teacher,omitempty"`
	StudentAttended bool      `json:"student_attended,omitempty"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
