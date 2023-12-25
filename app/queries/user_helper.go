package queries

import "github.com/icaksh/cripis/app/models"

func (q *UserQueries) GetUserRoles() ([]models.UserRoles, error) {
	users := []models.UserRoles{}
	query := `SELECT * FROM user_roles ORDER BY id`

	err := q.Select(&users, query)
	if err != nil {
		return users, err
	}

	return users, nil
}
