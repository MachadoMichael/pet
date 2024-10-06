package model

import (
	"github.com/oklog/ulid"
)

type SealUser struct {
	ID     ulid.ULID `json:"id"`
	UserID ulid.ULID `json:"user_id"`
	SealID ulid.ULID `json:"seal_id"`
}
