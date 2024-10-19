package model

import (
	"github.com/oklog/ulid"
)

type UserSeal struct {
	ID     ulid.ULID `json:"id"`
	UserID ulid.ULID `json:"user_id"`
	SealID ulid.ULID `json:"seal_id"`
}
