package dto

import (
	"github.com/oklog/ulid"
)

type NewPetDTO struct {
	CustomerID       ulid.ULID `json:"customer_id"`
	EmergencyContact string    `json:"emergency_contact"`
	Description      string    `json:"description"`
	Observation      string    `json:"observation"`
	Name             string    `json:"name"`
	Breed            string    `json:"breed"`
	Specie           string    `json:"specie"`
	Tag              string    `json:"tag"`
	Age              int       `json:"age"`
	Weight           int       `json:"weight"`
}
