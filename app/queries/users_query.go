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

func (q *UserQueries) GetUsers() ([]models.User, error) {
	users := []models.User{}
	query := `SELECT * FROM users`

	err := q.Get(&users, query)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (q *UserQueries) GetUser(id uuid.UUID) (models.User, error) {
	user := models.User{}
	query := `SELECT * FROM users WHERE id = $1`

	err := q.Get(&user, query, id)
	if err != nil {
		return user, err
	}
return user, nil
}

func (q *UserQueries) CreateUser(b *models.User) error {
	// Define query string.
	query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	// Send query to database.
	_, err := q.Exec(query, b.ID, b.CreatedAt, b.UpdatedAt, b.Email, b.Username, b.Password, b.Phone, b.Level, b.Status, b.Verified)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

func (q *UserQueries) CreateUserProfile(u *models.UserProfile) error {
	// Define query string.
	query := `INSERT INTO user_profile VALUES ($1, $2, $3, $4, $5)`

	// Send query to database.
	_, err := q.Exec(query, u.ID, u.CreatedAt, u.UpdatedAt, u.FirstName, u.LastName)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

func (q *UserQueries) CheckDuplicate(column string, value string) (bool) {
	users := models.User{}
	query := fmt.Sprintf(`SELECT * FROM users WHERE %s=$1`,column)
	
	err := q.Get(&users, query, value)
	fmt.Println(err)
	return err == nil
}