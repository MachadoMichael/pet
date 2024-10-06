package model

import (
	"github.com/oklog/ulid"
)

type Photo struct {
	ID     ulid.ULID `json:"id"`
	UserID ulid.ULID `json:"user_id"`
	Base64 string    `json:"image"`
}
