package database

import (
	"database/sql"

	"github.com/MachadoMichael/pet/infra/database"
	"github.com/MachadoMichael/pet/model"
	"github.com/oklog/ulid"
)

type CarerServiceRepo struct {
	DB        *sql.DB
	TableName string
}

func NewCarerServiceRepository(db *sql.DB) database.Repository[model.CarerService] {
	return &CarerServiceRepo{
		DB:        db,
		TableName: "carer_services",
	}
}

func (c *CarerServiceRepo) Create(model model.CarerService) (model.CarerService, error) {
	return model, nil
}

func (c *CarerServiceRepo) Read() ([]model.CarerService, error) {
	return []model.CarerService{}, nil
}

func (c *CarerServiceRepo) ReadOne(id ulid.ULID) (model.CarerService, error) {
	return model.CarerService{}, nil
}

func (c *CarerServiceRepo) Update(id ulid.ULID, model model.CarerService) error {
	return nil
}

func (c *CarerServiceRepo) Delete(id ulid.ULID) error {
	return nil
}

func (c *CarerServiceRepo) ReadQuery(query string, args ...any) ([]model.CarerService, error) {
	return []model.CarerService{}, nil
}

func (c *CarerServiceRepo) Query(query string, args ...any) (*sql.Rows, error) {
	return nil, nil
}
