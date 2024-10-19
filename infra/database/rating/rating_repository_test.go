package database

import (
	"testing"

	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestRatingRepo(t *testing.T) {
	repo := &RatingRepo{
		DB:        nil,
		TableName: "ratings",
	}

	assert.NotNil(t, repo)
}

func TestRatingRepo_Create(t *testing.T) {
	repo := &RatingRepo{
		DB:        nil,
		TableName: "ratings",
	}

	m, err := repo.Create(model.Rating{})
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, model.Rating{}, m)
}

func TestRatingRepo_Read(t *testing.T) {
	repo := &RatingRepo{
		DB:        nil,
		TableName: "ratings",
	}

	m, err := repo.Read()
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, []model.Rating{}, m)
}

func TestRatingRepo_ReadOne(t *testing.T) {
	repo := &RatingRepo{
		DB:        nil,
		TableName: "ratings",
	}

	m, err := repo.ReadOne(ulid.MustNew(ulid.Now(), nil))
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, model.Rating{}, m)
}

func TestRatingRepo_Update(t *testing.T) {
	repo := &RatingRepo{
		DB:        nil,
		TableName: "ratings",
	}

	err := repo.Update(ulid.MustNew(ulid.Now(), nil), model.Rating{})
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}

func TestRatingRepo_Delete(t *testing.T) {
	repo := &RatingRepo{
		DB:        nil,
		TableName: "ratings",
	}

	err := repo.Delete(ulid.MustNew(ulid.Now(), nil))
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}

func TestRatingRepo_ReadQuery(t *testing.T) {
	repo := &RatingRepo{
		DB:        nil,
		TableName: "ratings",
	}

	m, err := repo.ReadQuery("SELECT * FROM ratings")
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, []model.Rating{}, m)
}

func TestRatingRepo_Query(t *testing.T) {
	repo := &RatingRepo{
		DB:        nil,
		TableName: "ratings",
	}

	_, err := repo.Query("SELECT * FROM ratings")
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}
