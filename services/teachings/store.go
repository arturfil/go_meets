package teachings

import (
	"context"
	"database/sql"
	"errors"

	"github.com/arturfil/meetings_app_server/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) GetAllTeachings(userId string) ([]types.Teaching, error) {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	query := `
        SELECT id, teacher_id, subject_id, created_at, updated_at FROM teachings
        WHERE teacher_id = $1
    `

	rows, err := s.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	var teachings []types.Teaching

	for rows.Next() {
		var teaching types.Teaching
		err := rows.Scan(
			&teaching.ID,
			&teaching.TeacherId,
			&teaching.SubjectId,
			&teaching.CreatedAt,
			&teaching.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		teachings = append(teachings, teaching)
	}

	return teachings, nil
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

func (s *Store) GetSchedules(userId string) ([]types.Schedule, error) {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	query_count := `
        select count(*) from availability where user_id = $1;
    `

	var count int
	row := s.db.QueryRowContext(ctx, query_count, userId)
	err := row.Scan(&count)

	if count == 0 {
		return nil, errors.New("available times haven't been selected or user doesn't exist, please create new ones if possible")
	}

	query := `
        SELECT * FROM availability
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
        insert into availability (
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
        delete from availability where user_id = $1;
    `

	_, err := s.db.ExecContext(ctx, query, userId)
	if err != nil {
		return err
	}

	return nil
}
