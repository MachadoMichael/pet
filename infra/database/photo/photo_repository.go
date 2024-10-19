package database

import (
	"database/sql"

	"github.com/MachadoMichael/pet/infra/database"
	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
)

type PhotoRepo struct {
	DB        *sql.DB
	TableName string
}

func NewPhotoRepository(db *sql.DB) database.Repository[model.Photo] {
	return &PhotoRepo{
		DB:        db,
		TableName: "photos",
	}
}

func (c *PhotoRepo) Create(model model.Photo) (model.Photo, error) {
	return model, nil
}

func (c *PhotoRepo) Read() ([]model.Photo, error) {
	return []model.Photo{}, nil
}

func (c *PhotoRepo) ReadOne(id ulid.ULID) (model.Photo, error) {
	return model.Photo{}, nil
}

func (c *PhotoRepo) Update(id ulid.ULID, model model.Photo) error {
	return nil
}

func (c *PhotoRepo) Delete(id ulid.ULID) error {
	return nil
}

func (c *PhotoRepo) ReadQuery(query string, args ...any) ([]model.Photo, error) {
	return []model.Photo{}, nil
}

func (c *PhotoRepo) Query(query string, args ...any) (*sql.Rows, error) {
	return nil, nil
}
