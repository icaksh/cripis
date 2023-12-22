package queries

import "github.com/icaksh/cripis/app/models"

func (q *TrademarkQueries) GetTrademarkOwners() ([]models.TrademarkOwner, error) {
	result := []models.TrademarkOwner{}
	query := `SELECT * FROM trademark_owners`

	err := q.Get(&result, query)
	if err != nil {
		return result, err
	}

	return result, nil
}
