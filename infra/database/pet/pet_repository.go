package database

import (
	"database/sql"

	"github.com/MachadoMichael/pet/infra/database"
	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
)

type PetRepo struct {
	DB        *sql.DB
	TableName string
}

func NewPetRepository(db *sql.DB) database.Repository[model.Pet] {
	return &PetRepo{
		DB:        db,
		TableName: "pets",
	}
}

func (c *PetRepo) Create(model model.Pet) (model.Pet, error) {
	return model, nil
}

func (c *PetRepo) Read() ([]model.Pet, error) {
	return []model.Pet{}, nil
}

func (c *PetRepo) ReadOne(id ulid.ULID) (model.Pet, error) {
	return model.Pet{}, nil
}

func (c *PetRepo) Update(id ulid.ULID, model model.Pet) error {
	return nil
}

func (c *PetRepo) Delete(id ulid.ULID) error {
	return nil
}

func (c *PetRepo) ReadQuery(query string, args ...any) ([]model.Pet, error) {
	return []model.Pet{}, nil
}

func (c *PetRepo) Query(query string, args ...any) (*sql.Rows, error) {
	return nil, nil
}
