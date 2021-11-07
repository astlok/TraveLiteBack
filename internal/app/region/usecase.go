package region

import "travalite/internal/models"

type UseCase struct {
	repo Repo
}

func NewUseCase(repo Repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) GetRegions() ([]models.Region, error) {
	regions, err := u.repo.GetRegions()
	if err != nil {
		return nil, err
	}
	return regions, nil
}

func (u *UseCase) GetRegionInfo(id uint64) (models.Region, error) {
	region, err := u.repo.SelectRegionByID(id)
	if err != nil {
		return models.Region{}, err
	}
	return region, nil
}

