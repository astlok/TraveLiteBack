package profile

import (
	"database/sql"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"travalite/internal/models"
	customErrors "travalite/pkg/errors"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) Create(user models.User) (uint64, error) {
	var id uint64
	err := r.db.QueryRow(
		`INSERT INTO travelite.users (email, nickname, password)
	VALUES ($1, $2, $3) RETURNING id`,
		user.Email,
		user.Nickname,
		user.Password).Scan(&id)

	if err, ok := err.(*pq.Error); ok {
		if err.Code == pgerrcode.UniqueViolation {
			if err.Constraint == "users_email_key" {
				return 0, customErrors.DuplicateEmail
			}
			if err.Constraint == "users_nickname_key" {
				return 0, customErrors.DuplicateNickName
			}
		}
	}
	return id, nil
}

func (r *Repo) GetUserByEmailAndPass(email string, password string) (models.User, error) {
	user := models.User{}
	err := r.db.Get(&user, "SELECT * FROM travelite.users WHERE email = $1 AND password = $2", email, password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, customErrors.BadAuth
		}
		return user, err
	}
	return user, nil
}

func (r *Repo) GetUserByID(id uint64) (models.User, error) {
	user := models.User{}
	err := r.db.Get(&user, "SELECT * FROM travelite.users WHERE id=$1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, customErrors.UserNotFound
		}
		return user, err
	}
	return user, nil
}

func (r *Repo) ChangeProfile(u models.User) error {
	_, err := r.db.NamedExec(`UPDATE travelite.users SET email =:email, nickname=:nickname, password=:password, img=:img WHERE id=:id`, u)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code == pgerrcode.UniqueViolation {
				if err.Constraint == "users_email_key" {
					return customErrors.DuplicateEmail
				}
				if err.Constraint == "users_nickname_key" {
					return customErrors.DuplicateNickName
				}
			}
		}
		return err
	}

	return nil
}
