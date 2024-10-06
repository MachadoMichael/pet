package model

import (
	"github.com/oklog/ulid"
)

type CarerService struct {
	ID          ulid.ULID `json:"id"`
	CarerID     ulid.ULID `json:"carer_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
}
