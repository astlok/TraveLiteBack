package region

import (
	"github.com/jmoiron/sqlx"
	"travalite/internal/models"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) GetRegions() ([]models.Region, error) {
	var regions []models.Region
	err := r.db.Select(&regions, "SELECT * FROM travelite.region")
	if err != nil {
		return nil, err
	}
	return regions, err
}

func (r *Repo) SelectRegionByID(id uint64) (models.Region, error) {
	var region models.Region
	err := r.db.Get(&region, "SELECT * FROM travelite.region WHERE id = $1", id)
	if err != nil {
		return models.Region{}, err
	}
	return region, err
}
