package database

import (
	"testing"

	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestActionRepo(t *testing.T) {
	repo := &ActionRepo{
		DB:        nil,
		TableName: "actions",
	}

	assert.NotNil(t, repo)
}

func TestActionRepo_Create(t *testing.T) {
	repo := &ActionRepo{
		DB:        nil,
		TableName: "actions",
	}

	m, err := repo.Create(model.Action{})
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, model.Action{}, m)
}

func TestActionRepo_Read(t *testing.T) {
	repo := &ActionRepo{
		DB:        nil,
		TableName: "actions",
	}

	m, err := repo.Read()
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, []model.Action{}, m)
}

func TestActionRepo_ReadOne(t *testing.T) {
	repo := &ActionRepo{
		DB:        nil,
		TableName: "actions",
	}

	m, err := repo.ReadOne(ulid.MustNew(ulid.Now(), nil))
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, model.Action{}, m)
}

func TestActionRepo_Update(t *testing.T) {
	repo := &ActionRepo{
		DB:        nil,
		TableName: "actions",
	}

	err := repo.Update(ulid.MustNew(ulid.Now(), nil), model.Action{})
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}

func TestActionRepo_Delete(t *testing.T) {
	repo := &ActionRepo{
		DB:        nil,
		TableName: "actions",
	}

	err := repo.Delete(ulid.MustNew(ulid.Now(), nil))
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}

func TestActionRepo_ReadQuery(t *testing.T) {
	repo := &ActionRepo{
		DB:        nil,
		TableName: "actions",
	}

	m, err := repo.ReadQuery("SELECT * FROM actions")
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, []model.Action{}, m)
}

func TestActionRepo_Query(t *testing.T) {
	repo := &ActionRepo{
		DB:        nil,
		TableName: "actions",
	}

	_, err := repo.Query("SELECT * FROM actions")
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}
