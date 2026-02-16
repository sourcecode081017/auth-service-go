package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID         uuid.UUID
	FirstName    string
	LastName     string
	EmailAddress string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
