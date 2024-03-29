package models

import "github.com/google/uuid"

// Person is a entity that represents a person in all Domains
type Person struct {
	// ID is the identifier of the Entity, the ID is shared for all sub domains
	ID   uuid.UUID
	Name string
	Age  int
}
