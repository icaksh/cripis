package queries

import (
	"fmt"
	"github.com/icaksh/cripis/app/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserQueries struct {
	*sqlx.DB
	LogQueries
}

func (q *UserQueries) CreateUser(b *models.User) error {
	query := `INSERT INTO users(id, created_at, updated_at, email, first_name, last_name, password, phone, roles, status, verified) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := q.Exec(query, b.ID, b.CreatedAt, b.UpdatedAt, b.Email, b.FirstName, b.LastName, b.Password, b.Phone, b.Roles, b.Status, b.Verified)
	if err != nil {
		// Return only error.
		return err
	}

	return nil
}

func (q *UserQueries) GetUsers() ([]models.User, error) {
	users := []models.User{}
	query := `SELECT * FROM users ORDER BY users.created_at`

	err := q.Select(&users, query)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (q *UserQueries) GetUser(id uuid.UUID) (models.User, error) {
	user := models.User{}
	query := `SELECT * FROM users WHERE users.id = $1`

	err := q.Get(&user, query, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (q *UserQueries) CheckDuplicateUsers(column string, value string) (bool, error) {
	query := fmt.Sprintf(`SELECT EXISTS ( SELECT 1 FROM users WHERE users.%s=$1)`, column)
	var exists bool
	err := q.QueryRow(query, value).Scan(&exists)
	if err != nil {
		return false, nil
	}
	return exists, err
}
