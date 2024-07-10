package teachings

import (
	"context"
	"database/sql"

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

func (s *Store) CreateTeaching(userId, subjectId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

	query := `
        INSERT INTO teachings (
           teacher_id,
           subject_id,
           created_at,
           updated_at
        )
        VALUES ($1, $2)
    `

	_, err := s.db.ExecContext(
		ctx,
		query,
		userId,
		subjectId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetSchedule(userId string) (*types.Schedule, error) {
	ctx, cancel := context.WithTimeout(context.Background(), types.DBTimeout)
	defer cancel()

    query := `
        select user_id, start_time, end_time, created_at, updated_at from availability 
        where user_id = $1
    `

    var schedule types.Schedule

    row := s.db.QueryRowContext( ctx, query, userId)
    err := row.Scan(
        &schedule.UserId,
        &schedule.StartTime,
        &schedule.EndTime,
        &schedule.CreatedAt,
        &schedule.UpdatedAt,
    )
    if err != nil {
        return &types.Schedule{}, err
    }

    return &schedule, nil

}
