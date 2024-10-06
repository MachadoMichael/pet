package model

import (
	"github.com/oklog/ulid"
)

type Action struct {
	ID         ulid.ULID `json:"id"`
	CustomerID ulid.ULID `json:"customer_id"`
	PetID      ulid.ULID `json:"pet_id"`
	CarerID    ulid.ULID `json:"carer_id"`
	ServiceID  ulid.ULID `json:"service_id"`
	Duration   int       `json:"duration"` // minutes
}
