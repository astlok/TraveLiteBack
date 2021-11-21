package trek

import "travalite/internal/models"

type UseCase struct {
	repo Repo
}

func NewUseCase(repo Repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) CreateTrek(userID uint64, t models.Trek) (models.Trek, error) {
	trek, err := u.repo.CreateTrek(userID, t)

	if err != nil {
		return models.Trek{}, err
	}

	return trek, nil
}

func (u *UseCase) GetTrekInfo(ID uint64) (models.Trek, error) {
	trek, err := u.repo.SelectTrekById(ID)

	if err != nil {
		return models.Trek{}, nil
	}

	return trek, nil
}

func (u *UseCase) DeleteTrek(ID uint64) error {
	err := u.repo.DeleteTrek(ID)

	if err != nil {
		return err
	}

	return nil
}

func (u *UseCase) GetUsersTreks(ID uint64) ([]models.Trek, error) {
	treks, err := u.repo.GetAllUserTrek(ID)
	if err != nil {
		return nil, err
	}

	return treks, nil
}
