package model

import (
	"github.com/oklog/ulid"
)

type Pet struct {
	ID               ulid.ULID `json:"id"`
	CustomerID       ulid.ULID `json:"customer_id"`
	EmergencyContact string    `json:"emergency_contact"`
	Description      string    `json:"description"`
	Observation      string    `json:"observation"`
	Name             string    `json:"name"`
	Breed            string    `json:"breed"`
	Specie           string    `json:"specie"`
	Tag              string    `json:"tag"`
	Weight           float64   `json:"weight"`
	Age              int       `json:"age"`
}
