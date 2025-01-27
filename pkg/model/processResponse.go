package model

import (
	"github.com/google/uuid"
)

type ProcessResponse struct {
	ID uuid.UUID `json:"id"`
}