package trek

import (
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

//func (r *Repo) SearchTreks(keyword string) ([]models., error) {
//	var treks []models.//Trek
//	if keyword == "" {
//		return nil, nil
//	}
//	keyword += ":*"
//	if err := r.db.Select(&treks, /*Sql search*/, keyword); err != nil {
//		return nil, err
//	}
//	if len(treks) == 0 {
//		if err := r.db.Select(&treks, /*sql search*/, keyword); err != nil {
//			return nil, err
//		}
//	}
//	return treks, nil
//}
