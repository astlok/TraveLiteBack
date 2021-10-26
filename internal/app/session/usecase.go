package session

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"travalite/internal/models"
)

type UseCase struct {
	repo Repo
}

func NewUseCase(repo Repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) Create(user models.User) (models.Session, error) {
	var err error
	s := models.Session{UserID: user.ID}

	s.AuthToken, err = generateToken(user.Email)
	if err != nil {
		return models.Session{}, err
	}

	err = u.repo.DelSessionByUserID(s)
	if err != nil {
		return models.Session{}, err
	}

	err = u.repo.Create(s)
	if err != nil {
		return models.Session{}, err
	}

	return s, nil
}

func (u *UseCase) DelSession(session models.Session) error {
	err := u.repo.DelSessionByAuthToken(session)
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCase) Check(sessionID string) (models.Session, error) {
	s, err := u.repo.CheckSession(sessionID)
	if err != nil {
		return models.Session{}, err
	}

	return s, nil
}

func generateToken(email string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hasher := md5.New()
	hasher.Write(hash)

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
