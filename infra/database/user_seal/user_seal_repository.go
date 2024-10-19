package database

import (
	"database/sql"

	"github.com/MachadoMichael/pet/infra/database"
	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
)

type UserSealRepo struct {
	DB        *sql.DB
	TableName string
}

func NewUserSealRepository(db *sql.DB) database.Repository[model.UserSeal] {
	return &UserSealRepo{
		DB:        db,
		TableName: "user_seals",
	}
}

func (c *UserSealRepo) Create(model model.UserSeal) (model.UserSeal, error) {
	return model, nil
}

func (c *UserSealRepo) Read() ([]model.UserSeal, error) {
	return []model.UserSeal{}, nil
}

func (c *UserSealRepo) ReadOne(id ulid.ULID) (model.UserSeal, error) {
	return model.UserSeal{}, nil
}

func (c *UserSealRepo) Update(id ulid.ULID, model model.UserSeal) error {
	return nil
}

func (c *UserSealRepo) Delete(id ulid.ULID) error {
	return nil
}

func (c *UserSealRepo) ReadQuery(query string, args ...any) ([]model.UserSeal, error) {
	return []model.UserSeal{}, nil
}

func (c *UserSealRepo) Query(query string, args ...any) (*sql.Rows, error) {
	return nil, nil
}
