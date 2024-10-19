package database

import (
	"database/sql"

	"github.com/MachadoMichael/pet/infra/database"
	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
)

type RatingRepo struct {
	DB        *sql.DB
	TableName string
}

func NewRatingRepository(db *sql.DB) database.Repository[model.Rating] {
	return &RatingRepo{
		DB:        db,
		TableName: "ratings",
	}
}

func (c *RatingRepo) Create(model model.Rating) (model.Rating, error) {
	return model, nil
}

func (c *RatingRepo) Read() ([]model.Rating, error) {
	return []model.Rating{}, nil
}

func (c *RatingRepo) ReadOne(id ulid.ULID) (model.Rating, error) {
	return model.Rating{}, nil
}

func (c *RatingRepo) Update(id ulid.ULID, model model.Rating) error {
	return nil
}

func (c *RatingRepo) Delete(id ulid.ULID) error {
	return nil
}

func (c *RatingRepo) ReadQuery(query string, args ...any) ([]model.Rating, error) {
	return []model.Rating{}, nil
}

func (c *RatingRepo) Query(query string, args ...any) (*sql.Rows, error) {
	return nil, nil
}
