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

func (s *Store) GetAllMeetings() ([]types.MeetingResponse, error) {
    ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout) 
    defer cancel()

    query := `
       SELECT 
            m.id, 
            s.name AS "subject",
            u.first_name AS "student",
            t.first_name AS "teacher",
            m.student_attended, 
            m.start_time,
            m.end_time,
            m.created_at,
            m.updated_at
        FROM meetings m
        INNER JOIN "subjects" s ON s.id = m.subject_id
        LEFT JOIN "users" u ON u.id = m.student_id
        LEFT JOIN "users" t ON t.id = m.teacher_id 
        ; 
    `

    rows, err := s.db.QueryContext(ctx, query)
    if err != nil {
        return nil, err 
    }

    var meetings []types.MeetingResponse
    for rows.Next() {
        var meeting types.MeetingResponse
        err := rows.Scan(
            &meeting.ID,
            &meeting.Subject,
            &meeting.Student,
            &meeting.Teacher,
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
