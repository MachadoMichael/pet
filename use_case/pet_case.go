package use_case

import (
	"crypto/rand"

	"github.com/MachadoMichael/pet/dto"
	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
)

type petCase struct {
	repository database.Repository
	photoCase  *photoCase
}

func NewPetCase(repository database.Repository, photoCase *photoCase) *petCase {
	return &petCase{
		repository: repository,
		photoCase:  photoCase,
	}
}

func (p *petCase) Create(dto dto.NewPetDTO, photos [5]string) (model.Pet, error) {
	pet := model.Pet{
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

	err := pet.Save()
	if err != nil {
		return pet, err
	}

	for _, p := range photos {
		photo := model.Photo{
			ID:         ulid.MustNew(ulid.Now(), rand.Reader),
			ProviderID: pet.ID.String(),
			Base64:     p,
		}

		err = photo.Save()
		if err != nil {
			return pet, err
		}
	}

	return p.repository.Save(pet)
}
