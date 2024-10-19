package database

import (
	"database/sql"

	"github.com/MachadoMichael/pet/infra/database"
	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
)

type SealRepo struct {
	DB        *sql.DB
	TableName string
}

func NewSealRepository(db *sql.DB) database.Repository[model.Seal] {
	return &SealRepo{
		DB:        db,
		TableName: "seals",
	}
}

func (c *SealRepo) Create(model model.Seal) (model.Seal, error) {
	return model, nil
}

func (c *SealRepo) Read() ([]model.Seal, error) {
	return []model.Seal{}, nil
}

func (c *SealRepo) ReadOne(id ulid.ULID) (model.Seal, error) {
	return model.Seal{}, nil
}

func (c *SealRepo) Update(id ulid.ULID, model model.Seal) error {
	return nil
}

func (c *SealRepo) Delete(id ulid.ULID) error {
	return nil
}

func (c *SealRepo) ReadQuery(query string, args ...any) ([]model.Seal, error) {
	return []model.Seal{}, nil
}

func (c *SealRepo) Query(query string, args ...any) (*sql.Rows, error) {
	return nil, nil
}
