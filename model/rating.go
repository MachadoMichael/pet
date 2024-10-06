package model

import (
	"github.com/oklog/ulid"
)

type Rating struct {
	ID             ulid.ULID `json:"id"`
	CarerID        ulid.ULID `json:"carer_id"`
	CustomerID     ulid.ULID `json:"customer_id"`
	CarerServiceID ulid.ULID `json:"carer_service_id"`
	Rating         int       `json:"rating"`
}
