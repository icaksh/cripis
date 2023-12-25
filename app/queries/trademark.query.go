package queries

import (
	"github.com/google/uuid"
	"github.com/icaksh/cripis/app/models"
	"github.com/jmoiron/sqlx"
)

type TrademarkQueries struct {
	*sqlx.DB
	LogQueries
}

type GetTrademarkByUser struct {
	models.Trademark
	models.TrademarkRegistration
}

func (q *TrademarkQueries) CreateTrademark(v *models.Trademark) error {
	query := `INSERT INTO trademarks(id, 
                       created_at, 
                       updated_at, 
                       created_by, 
                       registration_number, 
                       trademark_name, 
                       trademark_class, 
                       owner_name, 
                       address, 
                       village, 
                       district, 
                       regency, 
                       province, 
                       file, 
                       image, 
                       status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`

	_, err := q.Exec(query, v.ID, v.CreatedAt, v.UpdatedAt, v.CreatedBy, v.RegisterNumber, v.TrademarkName, v.Class, v.OwnerName, v.Address, v.Village, v.District, v.Regency, v.Province, v.File, v.Image, v.Status)
	if err != nil {
		// Return only error.
		return err
	}

	return nil
}

func (q *TrademarkQueries) GetAllTrademarks() ([]models.Trademark, error) {
	result := []models.Trademark{}
	query := `SELECT * FROM trademarks ORDER BY trademarks.created_at`

	err := q.Select(&result, query)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (q *TrademarkQueries) GetAllTrademarksByName(name string) ([]models.Trademark, error) {
	result := []models.Trademark{}
	query := `SELECT * FROM trademarks WHERE trademarks.name=$1 ORDER BY trademarks.created_at`
	err := q.Select(&result, query, name)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (q *TrademarkQueries) GetTrademarksByUser(userId uuid.UUID) ([]models.Trademark, error) {
	result := []models.Trademark{}
	query := `SELECT * FROM trademarks WHERE trademarks.created_by=$1 ORDER BY trademarks.created_at ASC`
	err := q.Select(&result, query, userId)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (q *TrademarkQueries) GetTrademarkById(id uuid.UUID) ([]GetTrademarkByUser, error) {
	result := []GetTrademarkByUser{}
	query := `SELECT * FROM trademarks WHERE trademarks.id=$1 ORDER BY trademarks.created_at`
	err := q.Get(&result, query, id)
	if err != nil {
		return result, err
	}
	return result, nil
}
