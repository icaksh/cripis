package queries

import (
	"github.com/icaksh/cripis/app/models"
)

func (q *TrademarkQueries) GetClasses() ([]models.TrademarkClass, error) {
	res := []models.TrademarkClass{}
	query := `SELECT * FROM trademark_classes ORDER BY trademark_classes.id`

	err := q.Select(&res, query)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (q *TrademarkQueries) GetClass(id int) (models.TrademarkClass, error) {
	res := models.TrademarkClass{}
	query := `SELECT * FROM trademark_classes WHERE trademark_classes.id=$1`

	err := q.Get(&res, query, id)
	if err != nil {
		return res, err
	}
	return res, nil
}
