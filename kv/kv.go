package kv

import "errors"

var (
	// ErrNotFound is returned from Get when object with that key is not found.
	ErrNotFound = errors.New("not found")

	// ErrOldGen is returned from Set when object's generation is too old.
	ErrOldGen = errors.New("old gen")
)

type Key []byte

type Object struct {
	Value []byte
	Gen   uint64
}

type KV interface {
	// Get returns an object by a key.
	// ErrNotFound is returned when object is not found.
	Get(Key) (*Object, error)

	// Set sets an object with generation by key.
	// ErrOldGen is returned when object's generation is too old.
	Set(Key, *Object) error

	// Closes closes the KV storage and returns encountered error, if any.
	Close() error
}
