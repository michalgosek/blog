package models

import "github.com/google/uuid"

// Person describes basic person data.
type Person struct {
	ID   uuid.UUID
	Name string
	Age  int
}
