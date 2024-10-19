package database

import (
	"testing"

	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestCustomerRepo(t *testing.T) {
	repo := &CustomerRepo{
		DB:        nil,
		TableName: "customers",
	}

	assert.NotNil(t, repo)
}

func TestCustomerRepo_Create(t *testing.T) {
	repo := &CustomerRepo{
		DB:        nil,
		TableName: "customers",
	}

	m, err := repo.Create(model.Customer{})
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, model.Customer{}, m)
}

func TestCustomerRepo_Read(t *testing.T) {
	repo := &CustomerRepo{
		DB:        nil,
		TableName: "customers",
	}

	m, err := repo.Read()
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, []model.Customer{}, m)
}

func TestCustomerRepo_ReadOne(t *testing.T) {
	repo := &CustomerRepo{
		DB:        nil,
		TableName: "customers",
	}

	m, err := repo.ReadOne(ulid.MustNew(ulid.Now(), nil))
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, model.Customer{}, m)
}

func TestCustomerRepo_Update(t *testing.T) {
	repo := &CustomerRepo{
		DB:        nil,
		TableName: "customers",
	}

	err := repo.Update(ulid.MustNew(ulid.Now(), nil), model.Customer{})
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}

func TestCustomerRepo_Delete(t *testing.T) {
	repo := &CustomerRepo{
		DB:        nil,
		TableName: "customers",
	}

	err := repo.Delete(ulid.MustNew(ulid.Now(), nil))
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}

func TestCustomerRepo_ReadQuery(t *testing.T) {
	repo := &CustomerRepo{
		DB:        nil,
		TableName: "customers",
	}

	m, err := repo.ReadQuery("SELECT * FROM customers")
	assert.NotNil(t, repo)
	assert.NoError(t, err)
	assert.Equal(t, []model.Customer{}, m)
}

func TestCustomerRepo_Query(t *testing.T) {
	repo := &CustomerRepo{
		DB:        nil,
		TableName: "customers",
	}

	_, err := repo.Query("SELECT * FROM customers")
	assert.NotNil(t, repo)
	assert.NoError(t, err)
}
