package database

import (
	"database/sql"

	"github.com/oklog/ulid"
)

type Repository[T any] interface {
	Create(model T) error
	Read() ([]T, error)
	ReadOne(id ulid.ULID) (T, error)
	Update(id ulid.ULID, model T) error
	Delete(id ulid.ULID) error
	ReadQuery(query string, args ...any) ([]T, error)
	Query(query string, args ...any) (*sql.Rows, error)
}

type Repo[T any] struct {
	db        *sql.DB
	tableName string
}
