package profile

import (
	"travalite/internal/app/session"
	"travalite/internal/models"
)

type UseCase struct {
	userRepo    Repo
	sessionRepo session.Repo
}

func NewUseCase(userRepo Repo, sessionRepo session.Repo) *UseCase {
	return &UseCase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

func (u *UseCase) Create(user models.User) (uint64, error) {
	id, err := u.userRepo.Create(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UseCase) Auth(user models.User) (models.User, error) {
	var err error
	user, err = u.userRepo.GetUserByEmailAndPass(user.Email, user.Password)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *UseCase) GetUser(id uint64) (models.User, error) {
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return models.User{}, nil
	}
	user.ID = id
	user.Password = ""

	return user, nil
}

func (u *UseCase) ChangeProfile(user models.User) error {
	err := u.userRepo.ChangeProfile(user)

	if err != nil {
		return err
	}

	return nil
}
