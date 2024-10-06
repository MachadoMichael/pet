package model

import (
	"github.com/oklog/ulid"
)

type Seal struct {
	ID    ulid.ULID `json:"id"`
	Name  string    `json:"name"`
	Image string    `json:"image"`
}
