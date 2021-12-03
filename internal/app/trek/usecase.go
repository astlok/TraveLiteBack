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

func (u *UseCase) GetTreks() ([]models.Trek, error) {
	treks, err := u.repo.SelectTreks()
	if err != nil {
		return nil, err
	}

	return treks, nil
}

func (u *UseCase) ChangeTreks(ID uint64, t map[string]interface{}) (models.Trek, error) {
	err := u.repo.ChangeTrek(ID, t)

	if err != nil {
		return models.Trek{}, err
	}

	if things, ok := t["things"]; ok {
		err = u.repo.DeleteTrekThings(ID)
		if err != nil {
			return models.Trek{}, err
		}

		for _, thing := range things.([]interface{}) {
			err = u.repo.InsertTrekThing(ID, thing.(string))
			if err != nil {
				return models.Trek{}, err
			}
		}
	}

	if regionInterface, ok := t["region"]; ok {
		region := regionInterface.(string)
		err = u.repo.ChangeRegion(ID, region)
		if err != nil {
			return models.Trek{}, err
		}
	}

	trek, err := u.repo.SelectTrekById(ID)

	if err != nil {
		return models.Trek{}, err
	}

	return trek, nil
}

func (u *UseCase) RateTrek(trekID uint64, trekRate map[string]uint64) error {
	err := u.repo.UpdateTrekRate(trekID, trekRate)
	if err != nil {
		return err
	}

	return nil
}

func (u *UseCase) CreateComment(ID uint64, comment models.TrekComment) (models.TrekComment, error) {
	commentID, err := u.repo.InsertTrekComment(ID, comment)
	if err != nil {
		return models.TrekComment{}, err
	}

	err = u.repo.InsertTrekCommentPhotos(commentID, comment.Photo)
	if err != nil {
		return models.TrekComment{}, err
	}

	c, err := u.repo.GetComment(commentID)
	if err != nil {
		return models.TrekComment{}, err
	}

	return c, nil
}

func (u *UseCase) GetComments() ([]models.TrekComment, error) {
	c, err := u.repo.GetComments()
	if err != nil {
		return []models.TrekComment{}, err
	}

	return c, nil
}
