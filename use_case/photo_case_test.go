package use_case

import (
	"crypto/rand"
	"testing"

	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(base64 string, userID ulid.ULID) error {
	ID := ulid.MustNew(ulid.Now(), rand.Reader)
	args := m.Called(ID, base64, userID)
	return args.Error(0)
}

func TestPhotoCase_Create(t *testing.T) {

	mockRepo := new(MockRepository)
	photoUseCase := use_case.NewPhotoCase(mockRepo)

	userID := ulid.MustParse("user_id")
	photo := model.Photo{
		ID:     ulid.MustNew(ulid.Now(), rand.Reader),
		UserID: userID,
		Base64: "photo_base64",
	}

	mockRepo.On("Save", photo.ID, photo.Base64, photo.UserID).Return(nil)
	_, err := photoUseCase.Create(photo, userID)
	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Equal(t, "photo_base64", photo)

	mockRepo.AssertCalled(t, "Save", &photo)

}
