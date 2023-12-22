package queries

import "github.com/icaksh/cripis/app/models"

func (q *UserQueries) Auth(username string) (models.User, error) {
	users := models.User{}
	query := `SELECT * FROM users WHERE email=$1`

	err := q.Get(&users, query, username)
	if err != nil {
		return users, err
	}

	return users, nil
}
