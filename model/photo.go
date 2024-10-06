package model

import (
	"github.com/oklog/ulid"
)

type Photo struct {
	ID         ulid.ULID `json:"id"`
	ProviderID string    `json:"provider_id"`
	Image      string    `json:"image"`
}
