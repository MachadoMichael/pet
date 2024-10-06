package use_case_test

import (
	"errors"
	"testing"

	"crypto/rand"

	"github.com/MachadoMichael/pet/dto"
	"github.com/MachadoMichael/pet/model"
	"github.com/MachadoMichael/pet/use_case"
	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository mocks the database.Repository interface
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(pet *model.Pet) error {
	args := m.Called(pet)
	return args.Error(0)
}

func TestPetCase_Create(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	petUseCase := use_case.NewPetCase(mockRepo)

	// DTO with input data
	dto := dto.NewPetDTO{
		CustomerID:       ulid.MustNew(ulid.Now(), rand.Reader),
		EmergencyContact: "911",
		Description:      "Small dog",
		Observation:      "Very active",
		Name:             "Buddy",
		Breed:            "Poodle",
		Specie:           "Dog",
		Tag:              "001",
		Age:              2,
		Weight:           5.0,
	}

	photos := [5]string{"photo1.png", "photo2.png"}

	petID := ulid.MustNew(ulid.Now(), rand.Reader)
	pet := model.Pet{
		ID:               petID,
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

	// Mock Save method to return nil (success)
	mockRepo.On("Save", &pet).Return(nil)

	// Act
	createdPet, err := petUseCase.Create(dto, photos)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, pet.ID, createdPet.ID)
	assert.Equal(t, "Buddy", createdPet.Name)

	// Ensure that Save was called exactly once
	mockRepo.AssertCalled(t, "Save", &pet)

	// Optionally, you can test Save failure scenarios by returning an error
	mockRepo.On("Save", &pet).Return(errors.New("failed to save pet"))
	_, err = petUseCase.Create(dto, photos)
	assert.Error(t, err)
}
