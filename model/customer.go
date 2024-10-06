package model

import (
	"github.com/oklog/ulid"
)

type Customer struct {
	ID                 ulid.ULID `json:"id"`
	Name               string    `json:"name"`
	Cpf                string    `json:"cpf"`
	Phone              string    `json:"phone"`
	Email              string    `json:"email"`
	Address            string    `json:"address"`
	Photo              string    `json:"photo"`
	CardNumber         int       `json:"card_number"`
	SecurityCardNumber int       `json:"security_card_number"`
}
