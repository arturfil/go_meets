package meetings

import (
	"context"
	"database/sql"

	"github.com/arturfil/meetings_app_server/types"
)

type Store struct {
    db *sql.DB
}

func NewStore(db *sql.DB) *Store {
    return &Store{db: db}
}

func (s *Store) GetAllMeetings() ([]types.Meeting, error) {
    ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout) 
    defer cancel()

    query := ` SELECT * FROM meetings;`

    rows, err := s.db.QueryContext(ctx, query)
    if err != nil {
        return nil, err 
    }

    var meetings []types.Meeting
    for rows.Next() {
        var meeting types.Meeting
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

func (s *Store) GetMeetingByID(id string) (types.Meeting, error) {
    return types.Meeting{}, nil
}


func (s *Store) CreateMeeting(meeting types.Meeting) error {
    ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout) 
    defer cancel()

    query := `
        INSERT INTO meetings (
            subject_id,
            student_id,
            teacher_id,
            student_attended,
            start_time,
            end_time,
            created_at,
            updated_at
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
    `

    _, err := s.db.ExecContext(
        ctx,
        query,
        meeting.SubjectID,
        meeting.StudentID,
        meeting.TeacherID,
        meeting.StudentAttended,
        meeting.StartTime,
        meeting.EndTime,
        meeting.CreatedAt,
        meeting.UpdatedAt,
    )
    if err != nil {
        return err
    }
    
    return nil
}

func (s *Store) UpdateMeeting(meeting types.Meeting) error {
    return nil
}

func (s *Store) DeleteMeeting(id string) error {
    return nil
}
