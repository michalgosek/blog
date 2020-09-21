package utils

import (
	"github.com/google/uuid"
)

// UUIDGenerator returns new genrated id value.
func UUIDGenerator() uuid.UUID {
	return uuid.New()
}
