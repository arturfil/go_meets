package services

import (
	"context"
	"time"
)

type Meeting struct {
	ID              string    `json:"id"`
	SubjectID       time.Time `json:"dateandtime"`
	StudentID       string    `json:"studentid"`
	TeacherID       string    `json:"teacherid"`
	StudentAttended bool      `json:"student_attended"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (m *Meeting) GetAllMeetings() ([]Meeting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from meetings`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var meetings []Meeting
	for rows.Next() {
		var meeting Meeting
		err := rows.Scan(
			&meeting.ID,
			&meeting.SubjectID,
			&meeting.StudentID,
			&meeting.TeacherID,
			&meeting.StudentAttended,
			&meeting.StartTime,
			&meeting.EndTime,
			&meeting.CreatedAt,
			&meeting.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		meetings = append(meetings, meeting)
	}
	return meetings, nil
}
