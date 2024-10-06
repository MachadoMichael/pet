package model

import (
	"github.com/oklog/ulid"
)

type PhotoUser struct {
	ID      ulid.ULID `json:"id"`
	UserID  ulid.ULID `json:"user_id"`
	PhotoID ulid.ULID `json:"photo_id"`
}
