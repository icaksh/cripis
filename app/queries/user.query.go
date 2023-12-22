package queries

import (
	"fmt"
	"github.com/icaksh/cripis/app/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserQueries struct {
	*sqlx.DB
}

func (q *UserQueries) CreateUser(b *models.User) error {
	query := `INSERT INTO users(id, created_at, updated_at, email, first_name, last_name, password, phone, roles, status, verified) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := q.Exec(query, b.ID, b.CreatedAt, b.UpdatedAt, b.Email, b.FirstName, b.LastName, b.Password, b.Phone, b.Level, b.Status, b.Verified)
	if err != nil {
		// Return only error.
		return err
	}

	return nil
}

func (q *UserQueries) GetAllUsers() ([]models.User, error) {
	users := []models.User{}
	query := `SELECT * FROM users ORDER BY users.created_at`

	err := q.Get(&users, query)
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

func (q *UserQueries) CheckDuplicateUsers(column string, value string) bool {
	users := models.User{}
	query := fmt.Sprintf(`SELECT * FROM users WHERE users.%s=$1`, column)

	err := q.Get(&users, query, value)
	fmt.Println(err)
	return err == nil
}

func (q *UserQueries) CheckDuplicateProfile(column string, value uuid.UUID) (bool, error) {
	query := fmt.Sprintf(`SELECT EXISTS ( SELECT 1 FROM user_profile WHERE user_profile.%s=$1)`, column)
	var exists bool
	err := q.QueryRow(query, value).Scan(&exists)
	if err != nil {
		return false, nil
	}
	return exists, err
}
