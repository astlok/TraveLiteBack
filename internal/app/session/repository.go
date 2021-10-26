package session

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

func (r *Repo) DelSessionByAuthToken(s models.Session) error {
	_, err := r.db.NamedExec(`DELETE FROM travelite.sessions WHERE auth_token = :auth_token`, s)
	return err
}

func (r *Repo) DelSessionByUserID(s models.Session) error {
	_, err := r.db.NamedExec(`DELETE FROM travelite.sessions WHERE user_id = :user_id`, s)
	return err
}

func (r *Repo) Create(s models.Session) error {
	_, err := r.db.NamedExec(`INSERT INTO travelite.sessions(user_id, auth_token) VALUES (:user_id, :auth_token)`, s)
	return err
}

func (r *Repo) CheckSession(sessionID string) (models.Session, error) {
	var s models.Session
	err := r.db.Get(&s, "SELECT * FROM travelite.sessions WHERE auth_token=$1", sessionID)
	if err != nil {
		return models.Session{}, err
	}
	return s, nil
}
