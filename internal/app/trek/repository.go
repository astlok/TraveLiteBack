package trek

import (
	"errors"
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

	selectRatingByTrekId     = "SELECT avg(rating) FROM travelite.trek_rating WHERE trek_id = $1"
	selectRegionIDByName     = "SELECT id FROM travelite.region WHERE name = $1"
	updateRegionTrek         = "UPDATE travelite.trek SET region_id = $1 WHERE id = $2"
	insertOrUpdateTrekRating = "INSERT INTO travelite.trek_rating (user_id, trek_id, rating) " +
		"VALUES ($1, $2, $3) " +
		"ON CONFLICT (user_id, trek_id) DO UPDATE SET rating = $3"
	insertTrekComment        = "INSERT INTO travelite.comment (trek_id, user_id, description)  VALUES ($1, $2, $3) RETURNING id"
	insertTrekCommentsPhotos = "INSERT INTO travelite.comment_photo (comment_id, photo_url) VALUES ($1, $2)"
	selectTrekComment        = "SELECT comment.id, trek_id, description, nickname " +
		"FROM travelite.comment LEFT JOIN travelite.users ON comment.user_id = users.id " +
		"WHERE comment.id = $1"
	selectTrekComments = "SELECT comment.id, trek_id, description, nickname " +
		"FROM travelite.comment LEFT JOIN travelite.users ON comment.user_id = users.id"
	selectTrekCommentPhoto = "SELECT photo_url FROM travelite.comment_photo WHERE comment_id = $1"
	selectAllTreks = "SELECT trek.id, trek.name, difficult, days, description, " +
		"file, region.name AS region_name, " +
		"(SELECT COALESCE(AVG(rating), 0) FROM travelite.trek_rating) AS rating " +
		"FROM travelite.trek " +
		"LEFT JOIN travelite.region ON trek.region_id = region.id"
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

func (r *Repo) SelectTreks() ([]models.Trek, error) {
	var treks []models.Trek
	err := r.db.Select(&treks,
		selectAllTreks)

	if err != nil {
		return []models.Trek{}, err
	}

	for i, _ := range treks {
		err = r.db.Select(&treks[i].Things, selectThingsByTrekIdQuery, treks[i].ID)
		if err != nil {
			return []models.Trek{}, err
		}
	}

	return treks, nil
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

var (
	mutableTrekFields = map[string]string{
		"name":        "name = :name",
		"difficult":   "difficult = :difficult",
		"days":        "days = :days",
		"description": "description = :description",
		"file":        "file = :file",
	}
)

func makeTrekUpdateQuery(t map[string]interface{}) (string, bool) {
	if len(t) == 0 {
		return "", false
	}

	counter := 0

	trekUpdateQuery := "UPDATE travelite.trek SET\n"

	for key, _ := range t {
		if _, ok := mutableTrekFields[key]; !ok {
			continue
		}
		if counter != 0 {
			trekUpdateQuery += ",\n"
		}
		trekUpdateQuery += mutableTrekFields[key]
		counter++
	}

	if counter == 0 {
		return "", false
	}

	trekUpdateQuery += "\nWHERE id = :ID;"

	return trekUpdateQuery, true
}

func (r *Repo) InsertTrekThing(ID uint64, thing string) error {
	var thingID uint64

	err := r.db.Get(&thingID, "SELECT id FROM travelite.things WHERE name = $1;", thing)

	if err != nil && err.Error() != "sql: no rows in result set" {
		return err
	}
	err = nil
	if thingID == 0 {
		err = r.db.QueryRow(
			`INSERT INTO travelite.things (name) VALUES ($1) RETURNING id;`, thing).Scan(&thingID)
		if err != nil {
			return err
		}
	}

	err = r.db.QueryRow(`INSERT INTO travelite.trek_things (trek_id, thing_id) VALUES ($1, $2);`, ID, thingID).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteTrekThings(ID uint64) error {
	_, err := r.db.Exec("DELETE FROM travelite.trek_things WHERE trek_id = $1", ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) ChangeTrek(ID uint64, t map[string]interface{}) error {
	trekUpdateQuery, ok := makeTrekUpdateQuery(t)

	if ok {
		t["ID"] = ID
		_, err := r.db.NamedExec(trekUpdateQuery, t)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Repo) ChangeRegion(ID uint64, region string) error {
	var regionID uint64
	err := r.db.Get(&regionID, selectRegionIDByName, region)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(updateRegionTrek, regionID, ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) UpdateTrekRate(trekID uint64, trekRate map[string]uint64) error {
	userId, ok := trekRate["user_id"]
	if !ok {
		return errors.New("bad json")
	}
	rate, ok := trekRate["rate"]
	if !ok {
		return errors.New("bad json")
	}

	_, err := r.db.Exec(insertOrUpdateTrekRating, userId, trekID, rate)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) InsertTrekComment(ID uint64, comment models.TrekComment) (uint64, error) {
	var commentID uint64
	err := r.db.QueryRow(insertTrekComment, comment.TrekId, ID, comment.Description).Scan(&commentID)

	if err != nil {
		return 0, err
	}

	return commentID, nil
}

func (r *Repo) InsertTrekCommentPhotos(commentID uint64, urls []string) error {
	for _, url := range urls {
		_, err := r.db.Exec(insertTrekCommentsPhotos, commentID, url)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Repo) GetComment(ID uint64) (models.TrekComment, error) {
	var comment models.TrekComment
	err := r.db.Get(&comment, selectTrekComment, ID)
	if err != nil {
		return models.TrekComment{}, err
	}

	err = r.db.Select(&comment.Photo, selectTrekCommentPhoto, ID)
	if err != nil {
		return models.TrekComment{}, err
	}

	return comment, nil
}

func (r *Repo) GetComments() ([]models.TrekComment, error) {
	var comments []models.TrekComment
	err := r.db.Select(&comments, selectTrekComments)
	if err != nil {
		return []models.TrekComment{}, err
	}

	for i, _ := range comments {
		err = r.db.Select(&comments[i].Photo, selectTrekCommentPhoto, comments[i].ID)
		if err != nil {
			return []models.TrekComment{}, err
		}
	}
	return comments, nil
}