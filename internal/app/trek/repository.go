package trek

import (
	"github.com/jmoiron/sqlx"
	"travalite/internal/models"
)

const (
	selectTrekByIdQuery = "SELECT " +
		"trek.id, trek.name, difficult, days, description, file, region.name AS region_name " +
		"FROM " +
		"travelite.trek " +
		"LEFT JOIN " +
		"travelite.region " +
		"ON trek.region_id = region.id " +
		"WHERE trek.id = $1"

	selectThingsByTrekIdQuery = "SELECT travelite.things.name " +
		"FROM travelite.things " +
		"LEFT JOIN travelite.trek_things tt on things.id = tt.thing_id " +
		"WHERE tt.trek_id = $1"

	selectAllTrekByUserId = "SELECT t.id, t.name, difficult, days, description, file, r.name AS region_name " +
		"FROM " +
		"travelite.trek AS t " +
		"LEFT JOIN travelite.region AS r " +
		"ON t.region_id = r.id " +
		"WHERE t.user_id = $1"

	selectRatingByTrekId = "SELECT avg(rating) FROM travelite.trek_rating WHERE trek_id = $1"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) CreateTrek(userID uint64, t models.Trek) (models.Trek, error) {
	var ID uint64

	tx := r.db

	err := tx.QueryRow(
		`INSERT INTO travelite.trek (name, difficult, days, description, file, region_id, user_id)
	VALUES ($1, $2, $3, $4, $5, (SELECT id FROM travelite.region WHERE name = $6), $7) RETURNING id;`,
		t.Name,
		t.Difficult,
		t.Days,
		t.Description,
		t.File,
		t.Region,
		userID).Scan(&ID)
	if err != nil {
		return models.Trek{}, err
	}

	for _, thing := range t.Things {
		var thingID uint64

		err = tx.Get(&thingID, "SELECT id FROM travelite.things WHERE name = $1;", thing)

		if err != nil && err.Error() != "sql: no rows in result set" {
			return models.Trek{}, err
		}

		if thingID == 0 {
			err = tx.QueryRow(
				`INSERT INTO travelite.things (name) VALUES ($1) RETURNING id;`, thing).Scan(&thingID)
			if err != nil {
				return models.Trek{}, err
			}
		}

		err = tx.QueryRow(`INSERT INTO travelite.trek_things (trek_id, thing_id) VALUES ($1, $2);`, ID, thingID).Err()
		if err != nil {
			return models.Trek{}, err
		}
	}

	var trek models.Trek

	err = tx.Get(&trek, selectTrekByIdQuery, ID)
	if err != nil {
		return models.Trek{}, err
	}

	err = tx.Select(&trek.Things, selectThingsByTrekIdQuery, ID)
	if err != nil {
		return models.Trek{}, err
	}

	if err != nil {
		return models.Trek{}, err
	}

	return trek, nil
}

func (r *Repo) SelectTrekById(ID uint64) (models.Trek, error) {
	var trek models.Trek
	err := r.db.Get(&trek,
		selectTrekByIdQuery, ID)

	if err != nil {
		return models.Trek{}, err
	}

	err = r.db.Select(&trek.Things, selectThingsByTrekIdQuery, ID)
	if err != nil {
		return models.Trek{}, err
	}

	var rating *float64

	err = r.db.Get(&rating, selectRatingByTrekId, ID)

	if err != nil {
		return models.Trek{}, err
	}

	if rating == nil {
		trek.Rating = 0
	} else {
		trek.Rating = *rating
	}

	return trek, nil
}

func (r *Repo) DeleteTrek(ID uint64) error {
	_, err := r.db.Exec("DELETE FROM travelite.trek WHERE id = $1", ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetAllUserTrek(ID uint64) ([]models.Trek, error) {
	var treks []models.Trek

	err := r.db.Select(&treks, selectAllTrekByUserId, ID)

	if err != nil {
		return nil, err
	}

	for i, _ := range treks {
		err = r.db.Select(&treks[i].Things, selectThingsByTrekIdQuery, treks[i].ID)
		if err != nil {
			return nil, err
		}

		var rating *float64

		err = r.db.Get(&rating, selectRatingByTrekId, treks[i].ID)

		if err != nil {
			return nil, err
		}

		if rating == nil {
			treks[i].Rating = 0
		} else {
			treks[i].Rating = *rating
		}

	}

	return treks, nil

}
