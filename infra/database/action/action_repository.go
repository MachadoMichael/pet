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

func (c *ActionRepo) Create(model model.Action) (sql.Result, error) {
	rows, err := c.DB.Exec("INSERT INTO actions (id,customer_id, pet_id, carer_id, service_id, duration) VALUES ($1, $2, $3, $4, $5, $6)", model.ID, model.CustomerID, model.PetID, model.CarerID, model.ServiceID, model.Duration)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (c *ActionRepo) Read() (sql.Result, error) {
	actions, err := c.DB.Exec("SELECT * FROM actions")
	if err != nil {
		return nil, err
	}
	return actions, nil
}

func (c *ActionRepo) ReadOne(id ulid.ULID) (sql.Result, error) {
	action, err := c.DB.Exec("SELECT * FROM actions WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return action, nil
}

func (c *ActionRepo) Update(id ulid.ULID, model model.Action) error {
	return nil
}

func (c *ActionRepo) Delete(id ulid.ULID) (sql.Result, error) {
	rows, err := c.DB.Exec("DELETE FROM actions WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (c *ActionRepo) ReadQuery(query string, args ...any) ([]model.Action, error) {
	return []model.Action{}, nil
}

func (c *ActionRepo) Query(query string, args ...any) (*sql.Rows, error) {
	return nil, nil
}
