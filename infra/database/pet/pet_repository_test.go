package database

import (
	"testing"

	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestPetRepo(t *testing.T) {
	repo := &PetRepo{
		DB:        nil,
		TableName: "pets",
	}

	assert.NotNil(t, repo)
}

func TestPetRepo_Create(t *testing.T) {
	repo := &PetRepo{
		DB:        nil,
		TableName: "pets",
	}

	m, err := repo.Create(model.Pet{})
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, model.Pet{}, m)
}

func TestPetRepo_Read(t *testing.T) {
	repo := &PetRepo{
		DB:        nil,
		TableName: "pets",
	}

	m, err := repo.Read()
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, []model.Pet{}, m)
}

func TestPetRepo_ReadOne(t *testing.T) {
	repo := &PetRepo{
		DB:        nil,
		TableName: "pets",
	}

	m, err := repo.ReadOne(ulid.MustNew(ulid.Now(), nil))
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, model.Pet{}, m)
}

func TestPetRepo_Update(t *testing.T) {
	repo := &PetRepo{
		DB:        nil,
		TableName: "pets",
	}

	err := repo.Update(ulid.MustNew(ulid.Now(), nil), model.Pet{})
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}

func TestPetRepo_Delete(t *testing.T) {
	repo := &PetRepo{
		DB:        nil,
		TableName: "pets",
	}

	err := repo.Delete(ulid.MustNew(ulid.Now(), nil))
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}

func TestPetRepo_ReadQuery(t *testing.T) {
	repo := &PetRepo{
		DB:        nil,
		TableName: "pets",
	}

	m, err := repo.ReadQuery("SELECT * FROM pets")
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, []model.Pet{}, m)
}

func TestPetRepo_Query(t *testing.T) {
	repo := &PetRepo{
		DB:        nil,
		TableName: "pets",
	}

	_, err := repo.Query("SELECT * FROM pets")
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}
