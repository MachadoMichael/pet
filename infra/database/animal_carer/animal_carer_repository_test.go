package database

import (
	"testing"

	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestAnimalCarerRepo(t *testing.T) {
	repo := &AnimalCarerRepo{
		DB:        nil,
		TableName: "animal_carers",
	}

	assert.NotNil(t, repo)
}

func TestAnimalCarerRepo_Create(t *testing.T) {
	repo := &AnimalCarerRepo{
		DB:        nil,
		TableName: "animal_carers",
	}

	m, err := repo.Create(model.AnimalCarer{})
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, model.AnimalCarer{}, m)
}

func TestAnimalCarerRepo_Read(t *testing.T) {
	repo := &AnimalCarerRepo{
		DB:        nil,
		TableName: "animal_carers",
	}

	m, err := repo.Read()
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, []model.AnimalCarer{}, m)
}

func TestAnimalCarerRepo_ReadOne(t *testing.T) {
	repo := &AnimalCarerRepo{
		DB:        nil,
		TableName: "animal_carers",
	}

	m, err := repo.ReadOne(ulid.MustNew(ulid.Now(), nil))
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, model.AnimalCarer{}, m)
}

func TestAnimalCarerRepo_Update(t *testing.T) {
	repo := &AnimalCarerRepo{
		DB:        nil,
		TableName: "animal_carers",
	}

	err := repo.Update(ulid.MustNew(ulid.Now(), nil), model.AnimalCarer{})
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}

func TestAnimalCarerRepo_Delete(t *testing.T) {
	repo := &AnimalCarerRepo{
		DB:        nil,
		TableName: "animal_carers",
	}

	err := repo.Delete(ulid.MustNew(ulid.Now(), nil))
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}

func TestAnimalCarerRepo_ReadQuery(t *testing.T) {
	repo := &AnimalCarerRepo{
		DB:        nil,
		TableName: "animal_carers",
	}

	m, err := repo.ReadQuery("SELECT * FROM animal_carers")
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, []model.AnimalCarer{}, m)
}

func TestAnimalCarerRepo_Query(t *testing.T) {
	repo := &AnimalCarerRepo{
		DB:        nil,
		TableName: "animal_carers",
	}

	_, err := repo.Query("SELECT * FROM animal_carers")
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}
