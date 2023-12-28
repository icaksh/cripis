package queries

import (
	"fmt"
	"github.com/icaksh/cripis/app/models"
	"github.com/jmoiron/sqlx"
)

type UserQueries struct {
	*sqlx.DB
	LogQueries
}

func (q *UserQueries) CreateUser(b *models.User) error {
	query := `INSERT INTO users(
                  id, created_at, updated_at, 
                  email, first_name, last_name, 
                  password, roles, 
                  status, verified, card_number, 
                  dob, sex, address, 
                  village, district, regency,
                  province, postal_code) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)`

	_, err := q.Exec(query, b.ID, b.CreatedAt, b.UpdatedAt,
		b.Email, b.FirstName, b.LastName,
		b.Password, b.Roles,
		b.Status, b.Verified, b.CardNumber,
		b.DoB, b.Sex, b.Address,
		b.Village, b.District, b.Regency,
		b.Province, b.PostalCode)
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

func (q *UserQueries) GetUser(id string) (models.User, error) {
	user := models.User{}
	query := `SELECT * FROM users WHERE users.id = $1`

	err := q.Get(&user, query, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (q *UserQueries) UpdateUserProfile(v *models.User) error {
	query := `UPDATE users SET 
				email=$2, 
				first_name=$3,
				last_name=$4,
				card_number=$5,
				address=$6,
				dob=$7,
				sex=$8,
				village=$9,
				district=$10,
				regency=$11,
				province=$12,
				postal_code=$13,
                updated_at=$14
				WHERE id=$1`

	_, err := q.Exec(query, v.ID,
		v.Email, v.FirstName, v.LastName,
		v.CardNumber, v.Address, v.DoB,
		v.Sex, v.Village, v.District,
		v.Regency, v.Province, v.PostalCode,
		v.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (q *UserQueries) UpdateUserRoles(v *models.User) error {
	query := `UPDATE users SET 
				status=$2,
				roles=$3,
				verified=$4,
				updated_at=$5
				WHERE id=$1`

	_, err := q.Exec(query, v.ID,
		v.Status, v.Roles, v.Verified, v.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (q *UserQueries) UpdateUserPassword(v *models.User) error {
	query := `UPDATE users SET 
				password=$2,
				updated_at=$3
				WHERE id=$1`

	_, err := q.Exec(query, v.ID, v.Password, v.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (q *UserQueries) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id=$1`

	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
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
