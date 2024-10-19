package database

import (
	"database/sql"

	"github.com/MachadoMichael/pet/infra/database"
	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
)

type CustomerRepo struct {
	DB        *sql.DB
	TableName string
}

func NewCustomerRepository(db *sql.DB) database.Repository[model.Customer] {
	return &CustomerRepo{
		DB:        db,
		TableName: "customers",
	}
}

func (c *CustomerRepo) Create(model model.Customer) (model.Customer, error) {
	return model, nil
}

func (c *CustomerRepo) Read() ([]model.Customer, error) {
	return []model.Customer{}, nil
}

func (c *CustomerRepo) ReadOne(id ulid.ULID) (model.Customer, error) {
	return model.Customer{}, nil
}

func (c *CustomerRepo) Update(id ulid.ULID, model model.Customer) error {
	return nil
}

func (c *CustomerRepo) Delete(id ulid.ULID) error {
	return nil
}

func (c *CustomerRepo) ReadQuery(query string, args ...any) ([]model.Customer, error) {
	return []model.Customer{}, nil
}

func (c *CustomerRepo) Query(query string, args ...any) (*sql.Rows, error) {
	return nil, nil
}
