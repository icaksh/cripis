package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type LogQueries struct {
	*sqlx.DB
}

func (q *LogQueries) CreateLog(userId uuid.UUID, action string) {
	query := `INSERT INTO logs (user_id, action) VALUES ($1, $2)`

	_, err := q.Exec(query, userId, action)
	if err != nil {
		return
	}
	return
}
