package database

import (
	"database/sql"

	"github.com/MachadoMichael/pet/infra/database"
	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
)

type ActionRepo struct {
	DB        *sql.DB
	TableName string
}

func NewActionRepository(db *sql.DB) database.Repository[model.Action] {
	return &ActionRepo{
		DB:        db,
		TableName: "actions",
	}
}

func (c *ActionRepo) Create(model model.Action) (model.Action, error) {
	return model, nil
}

func (c *ActionRepo) Read() ([]model.Action, error) {
	return []model.Action{}, nil
}

func (c *ActionRepo) ReadOne(id ulid.ULID) (model.Action, error) {
	return model.Action{}, nil
}

func (c *ActionRepo) Update(id ulid.ULID, model model.Action) error {
	return nil
}

func (c *ActionRepo) Delete(id ulid.ULID) error {
	return nil
}

func (c *ActionRepo) ReadQuery(query string, args ...any) ([]model.Action, error) {
	return []model.Action{}, nil
}

func (c *ActionRepo) Query(query string, args ...any) (*sql.Rows, error) {
	return nil, nil
}
