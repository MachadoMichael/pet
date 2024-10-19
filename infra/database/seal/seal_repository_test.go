package database

import (
	"testing"

	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestSealRepo(t *testing.T) {
	repo := &SealRepo{
		DB:        nil,
		TableName: "seals",
	}

	assert.NotNil(t, repo)
}

func TestSealRepo_Create(t *testing.T) {
	repo := &SealRepo{
		DB:        nil,
		TableName: "seals",
	}

	m, err := repo.Create(model.Seal{})
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, model.Seal{}, m)
}

func TestSealRepo_Read(t *testing.T) {
	repo := &SealRepo{
		DB:        nil,
		TableName: "seals",
	}

	m, err := repo.Read()
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, []model.Seal{}, m)
}

func TestSealRepo_ReadOne(t *testing.T) {
	repo := &SealRepo{
		DB:        nil,
		TableName: "seals",
	}

	m, err := repo.ReadOne(ulid.MustNew(ulid.Now(), nil))
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, model.Seal{}, m)
}

func TestSealRepo_Update(t *testing.T) {
	repo := &SealRepo{
		DB:        nil,
		TableName: "seals",
	}

	err := repo.Update(ulid.MustNew(ulid.Now(), nil), model.Seal{})
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}

func TestSealRepo_Delete(t *testing.T) {
	repo := &SealRepo{
		DB:        nil,
		TableName: "seals",
	}

	err := repo.Delete(ulid.MustNew(ulid.Now(), nil))
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}

func TestSealRepo_ReadQuery(t *testing.T) {
	repo := &SealRepo{
		DB:        nil,
		TableName: "seals",
	}

	m, err := repo.ReadQuery("SELECT * FROM seals")
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, []model.Seal{}, m)
}

func TestSealRepo_Query(t *testing.T) {
	repo := &SealRepo{
		DB:        nil,
		TableName: "seals",
	}

	_, err := repo.Query("SELECT * FROM seals")
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}
