package database

import (
	"database/sql"

	"github.com/MachadoMichael/pet/infra/database"
	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
)

type AnimalCarerRepo struct {
	DB        *sql.DB
	TableName string
}

func NewAnimalCarerRepository(db *sql.DB) database.Repository[model.AnimalCarer] {
	return &AnimalCarerRepo{
		DB:        db,
		TableName: "animal_carers",
	}
}

func (c *AnimalCarerRepo) Create(model model.AnimalCarer) (model.AnimalCarer, error) {
	return model, nil
}

func (c *AnimalCarerRepo) Read() ([]model.AnimalCarer, error) {
	return []model.AnimalCarer{}, nil
}

func (c *AnimalCarerRepo) ReadOne(id ulid.ULID) (model.AnimalCarer, error) {
	return model.AnimalCarer{}, nil
}

func (c *AnimalCarerRepo) Update(id ulid.ULID, model model.AnimalCarer) error {
	return nil
}

func (c *AnimalCarerRepo) Delete(id ulid.ULID) error {
	return nil
}

func (c *AnimalCarerRepo) ReadQuery(query string, args ...any) ([]model.AnimalCarer, error) {
	return []model.AnimalCarer{}, nil
}

func (c *AnimalCarerRepo) Query(query string, args ...any) (*sql.Rows, error) {
	return nil, nil
}
