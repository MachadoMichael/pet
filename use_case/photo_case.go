package use_case

import (
	"errors"

	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
)

type photoCase struct {
	repository database.Repository
}

func NewPhotoCase(repository database.Repository) *photoCase {
	return &photoCase{
		repository: repository,
	}
}

func (p *photoCase) Create(photo model.Photo, userID ulid.ULID) error {
	photos, err := p.repository.FindByUserID(userID)
	if err != nil {
		return err
	}

	if len(photos) >= 5 {
		return errors.New("user already has 5 photos")
	}

	return p.repository.Save(photo.ID, photo.Base64, userID)
}
