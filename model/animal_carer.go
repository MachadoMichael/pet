package model

import (
	"github.com/oklog/ulid"
)

type AnimalCarer struct {
	ID               ulid.ULID `json:"id"`
	Name             string    `json:"name"`
	Address          string    `json:"address"`
	Phone            string    `json:"phone"`
	Email            string    `json:"email"`
	Description      string    `json:"description"`
	Tags             string    `json:"tags"`
	ProofOfResidence string    `json:"proof_of_residence"`
	PixKey           string    `json:"pix_key"`
	Age              int       `json:"age"`
}
