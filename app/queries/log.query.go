package queries

import (
	"github.com/google/uuid"
	"github.com/icaksh/cripis/app/models"
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

func (q *LogQueries) GetLogs() ([]models.Log, error) {
	result := []models.Log{}
	query := `SELECT logs.id, logs.timestamp, logs.user_id, logs.action, users.email, CONCAT( users.first_name, ' ',  users.last_name) as user_name FROM logs INNER JOIN users on users.id = logs.user_id ORDER BY logs.timestamp DESC`

	err := q.Select(&result, query)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (q *LogQueries) GetLogsTrademarksByYears(year int) ([]models.LogYear, error) {
	result := []models.LogYear{}
	query := `select extract(month from created_at) as month, 
       		  count(*) as count, 
			  COUNT(CASE WHEN status = 2 THEN 1 END) AS approved_count,
  			  COUNT(CASE WHEN status = 4 THEN 1 END) AS rejected_count
			  from trademarks 
			  where extract(year from created_at) = $1 group by month `

	err := q.Select(&result, query, year)
	if err != nil {
		return result, err
	}

	return result, nil
}
func (q *LogQueries) GetLogsTrademarksByMonth(year int, month int) ([]models.LogMonth, error) {
	result := []models.LogMonth{}
	query := `select
    			extract(day from created_at) as day, 
       		  count(*) as count,
  			  COUNT(CASE WHEN status = 2 THEN 1 END) AS approved_count,
  			  COUNT(CASE WHEN status = 4 THEN 1 END) AS rejected_count
			  from trademarks 
			  where extract(year from created_at) = $1 and
			  extract(month from created_at) = $2 group by day`

	err := q.Select(&result, query, year, month)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (q *LogQueries) GetLogsLoginByYear(year int) ([]models.LogUserYear, error) {
	result := []models.LogUserYear{}
	query := `select extract(month from timestamp) as month, 
       		  count(*) as count
			  from logs 
			  where extract(year from timestamp) = $1 
			  and action ilike '%login%'
			  group by month `

	err := q.Select(&result, query, year)
	if err != nil {
		return result, err
	}

	return result, nil
}
func (q *LogQueries) GetLogsLoginByMonth(year int, month int) ([]models.LogUserMonth, error) {
	result := []models.LogUserMonth{}
	query := `select
    			extract(day from timestamp) as day, 
       		  count(*) as count
			  from logs 
			  where extract(year from timestamp) = $1
			and
			  extract(month from timestamp) = $2 
			  and action ilike '%login%'
			  group by day`

	err := q.Select(&result, query, year, month)
	if err != nil {
		return result, err
	}

	return result, nil
}
