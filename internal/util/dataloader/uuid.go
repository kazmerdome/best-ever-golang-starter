package dataloader

import "github.com/google/uuid"

// UuidKey implements the Key interface for a uuid.UUID
type UuidKey uuid.UUID

// String is an identity method. Used to implement String interface
func (r UuidKey) String() string {
	return uuid.UUID(r).String()
}

// // String is an identity method. Used to implement Key Raw
func (r UuidKey) Raw() interface{} { return uuid.UUID(r) }
