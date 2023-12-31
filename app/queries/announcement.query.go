package queries

import (
	"github.com/icaksh/cripis/app/models"
	"github.com/jmoiron/sqlx"
	"time"
)

type AnnouncementQueries struct {
	*sqlx.DB
}

func (q *AnnouncementQueries) CreateAnnouncement(v *models.Announcement) error {
	query := `INSERT INTO announcements(created_at, updated_at, created_by, title, description, image) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := q.Exec(query, time.Now(), time.Now(), v.CreatedBy, v.Title, v.Description, v.Image)
	if err != nil {
		return err
	}
	return nil
}

func (q *AnnouncementQueries) GetAnnouncements() ([]models.Announcement, error) {
	result := []models.Announcement{}
	query := `SELECT * FROM announcements ORDER BY announcements.created_at DESC LIMIT 3`

	err := q.Select(&result, query)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (q *AnnouncementQueries) GetAllAnnouncements() ([]models.Announcement, error) {
	result := []models.Announcement{}
	query := `SELECT * FROM announcements ORDER BY announcements.created_at`

	err := q.Select(&result, query)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (q *AnnouncementQueries) GetAnnouncement(userId int) (models.Announcement, error) {
	result := models.Announcement{}
	query := `SELECT * FROM announcements WHERE announcements.id=$1`

	err := q.Get(&result, query, userId)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (q *AnnouncementQueries) UpdateAnnouncement(v *models.Announcement) error {
	query := `UPDATE announcements SET updated_at=$2, title=$3, description=$4, image=$5 WHERE id=$1`

	_, err := q.Exec(query, v.Id, time.Now(), v.Title, v.Description, v.Image)
	if err != nil {
		return err
	}
	return nil
}

func (q *AnnouncementQueries) DeleteAnnouncement(id int) error {
	query := `DELETE FROM announcements WHERE id=$1`

	_, err := q.Exec(query, id)
	if err != nil {
		// Return only error.
		return err
	}

	return nil
}
