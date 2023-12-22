package queries

import (
	"github.com/google/uuid"
	"github.com/icaksh/cripis/app/models"
)

func (q *UserQueries) CreateUserProfile(u *models.UserProfile) error {
	query := `INSERT INTO user_profiles(user_id, created_at, updated_at, card_number, dob, sex, address, village, district, regency, province, postal_code) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := q.Exec(query, u.ID, u.CreatedAt, u.UpdatedAt, u.CardNumber, u.DoB, u.Sex, u.Address, u.Village, u.District, u.Regency, u.Province, u.PostalCode)
	if err != nil {
		return err
	}
	return nil
}

func (q *UserQueries) GetAllUserProfile(id uuid.UUID) (models.UserProfile, error) {
	profile := models.UserProfile{}
	query := `SELECT * FROM user_profiles ORDER BY user_profiles.created_at`
	err := q.Get(&profile, query, id)
	if err != nil {
		return profile, err
	}
	return profile, nil
}

func (q *UserQueries) GetUserProfile(id uuid.UUID) (models.UserProfile, error) {
	profile := models.UserProfile{}
	query := `SELECT * FROM user_profiles WHERE user_profiles.user_id=$1`
	err := q.Get(&profile, query, id)
	if err != nil {
		return profile, err
	}
	return profile, nil
}

func (q *UserQueries) DeleteUserProfile(id uuid.UUID) error {
	query := `DELETE FROM user_profiles WHERE user_id=$1`
	_, err := q.Exec(query, id)
	if err != nil {
		// Return only error.
		return err
	}
	return nil
}
