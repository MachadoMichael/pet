package database

import (
	"testing"

	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestPhotoRepo(t *testing.T) {
	repo := &PhotoRepo{
		DB:        nil,
		TableName: "photos",
	}

	assert.NotNil(t, repo)
}

func TestPhotoRepo_Create(t *testing.T) {
	repo := &PhotoRepo{
		DB:        nil,
		TableName: "photos",
	}

	m, err := repo.Create(model.Photo{})
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, model.Photo{}, m)
}

func TestPhotoRepo_Read(t *testing.T) {
	repo := &PhotoRepo{
		DB:        nil,
		TableName: "photos",
	}

	m, err := repo.Read()
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, []model.Photo{}, m)
}

func TestPhotoRepo_ReadOne(t *testing.T) {
	repo := &PhotoRepo{
		DB:        nil,
		TableName: "photos",
	}

	m, err := repo.ReadOne(ulid.MustNew(ulid.Now(), nil))
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, model.Photo{}, m)
}

func TestPhotoRepo_Update(t *testing.T) {
	repo := &PhotoRepo{
		DB:        nil,
		TableName: "photos",
	}

	err := repo.Update(ulid.MustNew(ulid.Now(), nil), model.Photo{})
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}

func TestPhotoRepo_Delete(t *testing.T) {
	repo := &PhotoRepo{
		DB:        nil,
		TableName: "photos",
	}

	err := repo.Delete(ulid.MustNew(ulid.Now(), nil))
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}

func TestPhotoRepo_ReadQuery(t *testing.T) {
	repo := &PhotoRepo{
		DB:        nil,
		TableName: "photos",
	}

	m, err := repo.ReadQuery("SELECT * FROM photos")
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, []model.Photo{}, m)
}

func TestPhotoRepo_Query(t *testing.T) {
	repo := &PhotoRepo{
		DB:        nil,
		TableName: "photos",
	}

	_, err := repo.Query("SELECT * FROM photos")
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}
