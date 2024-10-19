package database

import (
	"testing"

	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestUserSealRepo(t *testing.T) {
	repo := &UserSealRepo{
		DB:        nil,
		TableName: "user_seals",
	}

	assert.NotNil(t, repo)
}

func TestUserSealRepo_Create(t *testing.T) {
	repo := &UserSealRepo{
		DB:        nil,
		TableName: "user_seals",
	}

	m, err := repo.Create(model.UserSeal{})
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, model.UserSeal{}, m)
}

func TestUserSealRepo_Read(t *testing.T) {
	repo := &UserSealRepo{
		DB:        nil,
		TableName: "user_seals",
	}

	m, err := repo.Read()
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, []model.UserSeal{}, m)
}

func TestUserSealRepo_ReadOne(t *testing.T) {
	repo := &UserSealRepo{
		DB:        nil,
		TableName: "user_seals",
	}

	m, err := repo.ReadOne(ulid.MustNew(ulid.Now(), nil))
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, model.UserSeal{}, m)
}

func TestUserSealRepo_Update(t *testing.T) {
	repo := &UserSealRepo{
		DB:        nil,
		TableName: "user_seals",
	}

	err := repo.Update(ulid.MustNew(ulid.Now(), nil), model.UserSeal{})
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}

func TestUserSealRepo_Delete(t *testing.T) {
	repo := &UserSealRepo{
		DB:        nil,
		TableName: "user_seals",
	}

	err := repo.Delete(ulid.MustNew(ulid.Now(), nil))
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}

func TestUserSealRepo_ReadQuery(t *testing.T) {
	repo := &UserSealRepo{
		DB:        nil,
		TableName: "user_seals",
	}

	m, err := repo.ReadQuery("SELECT * FROM user_seals")
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, []model.UserSeal{}, m)
}

func TestUserSealRepo_Query(t *testing.T) {
	repo := &UserSealRepo{
		DB:        nil,
		TableName: "user_seals",
	}

	_, err := repo.Query("SELECT * FROM user_seals")
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}
