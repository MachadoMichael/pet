package database

import (
	"testing"

	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestCarerServiceRepo(t *testing.T) {
	repo := &CarerServiceRepo{
		DB:        nil,
		TableName: "carer_services",
	}

	assert.NotNil(t, repo)
}

func TestCarerServiceRepo_Create(t *testing.T) {
	repo := &CarerServiceRepo{
		DB:        nil,
		TableName: "carer_services",
	}

	m, err := repo.Create(model.CarerService{})
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, model.CarerService{}, m)
}

func TestCarerServiceRepo_Read(t *testing.T) {
	repo := &CarerServiceRepo{
		DB:        nil,
		TableName: "carer_services",
	}

	m, err := repo.Read()
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, []model.CarerService{}, m)
}

func TestCarerServiceRepo_ReadOne(t *testing.T) {
	repo := &CarerServiceRepo{
		DB:        nil,
		TableName: "carer_services",
	}

	m, err := repo.ReadOne(ulid.MustNew(ulid.Now(), nil))
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, model.CarerService{}, m)
}

func TestCarerServiceRepo_Update(t *testing.T) {
	repo := &CarerServiceRepo{
		DB:        nil,
		TableName: "carer_services",
	}

	err := repo.Update(ulid.MustNew(ulid.Now(), nil), model.CarerService{})
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}

func TestCarerServiceRepo_Delete(t *testing.T) {
	repo := &CarerServiceRepo{
		DB:        nil,
		TableName: "carer_services",
	}

	err := repo.Delete(ulid.MustNew(ulid.Now(), nil))
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}

func TestCarerServiceRepo_ReadQuery(t *testing.T) {
	repo := &CarerServiceRepo{
		DB:        nil,
		TableName: "carer_services",
	}

	m, err := repo.ReadQuery("SELECT * FROM carer_services")
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, []model.CarerService{}, m)
}

func TestCarerServiceRepo_Query(t *testing.T) {
	repo := &CarerServiceRepo{
		DB:        nil,
		TableName: "carer_services",
	}

	_, err := repo.Query("SELECT * FROM carer_services")
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}
