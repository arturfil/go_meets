package teachings

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/arturfil/meetings_app_server/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) GetAllTeachings(userId string) ([]types.SubjectResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	query := `
        SELECT t.id, s.name, c.name, s.description, s.created_at, s.updated_at FROM teachings t
        JOIN users u ON u.id = t.teacher_id
        LEFT JOIN subjects s ON t.subject_id = s.id
        JOIN subject_categories c ON c.id = s.category_id
        WHERE u.id = $1;
    `

	rows, err := s.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	var subjects []types.SubjectResponse

	for rows.Next() {
		var subject types.SubjectResponse
		err := rows.Scan(
			&subject.ID,
			&subject.Name,
			&subject.Category,
			&subject.Description,
			&subject.CreatedAt,
			&subject.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		subjects = append(subjects, subject)
	}

	return subjects, nil
}

func (s *Store) CreateTeaching(teaching types.TeachingSubmission) error {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	query := `
        INSERT INTO teachings (
           teacher_id,
           subject_id
        )
        VALUES ($1, $2)
    `

	_, err := s.db.ExecContext(
		ctx,
		query,
		teaching.TeacherID,
		teaching.SubjectID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteTeaching(teachingId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

    query := `
        DELETE FROM teachings WHERE id = $1;
    `

    res, err := s.db.ExecContext(ctx, query, teachingId)
    if err != nil {
        return err
    }
    
    fmt.Println("result", res)
    return nil
}

func (s *Store) GetSchedules(userId string) ([]types.Schedule, error) {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	query_count := `
        SELECT count(*) FROM schedules WHERE user_id = $1;
    `

	var count int
	row := s.db.QueryRowContext(ctx, query_count, userId)
	err := row.Scan(&count)

	if count == 0 {
		return nil, errors.New("available times haven't been selected or user doesn't exist, please create new ones if possible")
	}

	query := `
        SELECT * FROM schedules
        WHERE user_id = $1;
    `

	rows, err := s.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	var schedules []types.Schedule

	for rows.Next() {
		var schedule types.Schedule
		err := rows.Scan(
			&schedule.UserId,
			&schedule.StartTime,
			&schedule.EndTime,
			&schedule.Day,
			&schedule.CreatedAt,
			&schedule.UpdatedAt,
		)
        if err != nil {
            return nil, err
        }

		schedules = append(schedules, schedule)
	}

	return schedules, nil
}

func (s *Store) CreateSchedule(schedule types.Schedule) error {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	query := `
        insert into schedules (
            user_id,
            start_time,
            end_time,
            day
        )
        values ($1, $2, $3, $4);
    `

	_, err := s.db.ExecContext(
		ctx,
		query,
		schedule.UserId,
		schedule.StartTime,
		schedule.EndTime,
		schedule.Day,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteSchedule(userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	query := `
        DELETE FROM schedules WHERE user_id = $1;
    `

	_, err := s.db.ExecContext(ctx, query, userId)
	if err != nil {
		return err
	}

	return nil
}
