package use_case

import (
	"crypto/rand"

	"github.com/MachadoMichael/pet/dto"
	"github.com/MachadoMichael/pet/infra/database"
	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
)

type petCase struct {
	repository database.Repository[model.Pet]
	photoCase  *photoCase
}

func NewPetCase(repo database.Repository[model.Pet], photoCase *photoCase) *petCase {
	return &petCase{
		repository: repo,
		photoCase:  photoCase,
	}
}

func (p *petCase) Create(dto dto.NewPetDTO, photos [5]string) (model.Pet, error) {
	petModel := model.Pet{
		ID:               ulid.MustNew(ulid.Now(), rand.Reader),
		CustomerID:       dto.CustomerID,
		EmergencyContact: dto.EmergencyContact,
		Description:      dto.Description,
		Observation:      dto.Observation,
		Name:             dto.Name,
		Breed:            dto.Breed,
		Specie:           dto.Specie,
		Tag:              dto.Tag,
		Age:              dto.Age,
		Weight:           dto.Weight,
	}

	for _, p := range photos {
		photoModel := model.Photo{
			ID:     ulid.MustNew(ulid.Now(), rand.Reader),
			UserID: petModel.CustomerID,
			Base64: p,
		}

		photo, err = photoCase.Create(p, petModel.CustomerID)
		if err != nil {
			return petModel, err
		}
	}

	return p.repository.Create(petModel)
}
