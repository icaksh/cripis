package queries

import (
	"github.com/icaksh/cripis/app/models"
	"github.com/jmoiron/sqlx"
)

type TrademarkHelpersQueries struct {
	*sqlx.DB
}

func (q *TrademarkHelpersQueries) GetClasses() ([]models.TrademarkClass, error) {
	res := []models.TrademarkClass{}
	query := `SELECT * FROM trademark_classes`

	err := q.Select(&res, query)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (q *TrademarkHelpersQueries) GetClass(id int) (models.TrademarkClass, error) {
	res := models.TrademarkClass{}
	query := `SELECT * FROM trademark_classes WHERE trademark.id=$1`

	err := q.Get(&res, query, id)
	if err != nil {
		return res, err
	}
	return res, nil
}
