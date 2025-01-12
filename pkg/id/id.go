package id

import "github.com/google/uuid"

type ID = uuid.UUID

func New() ID {
	return ID(uuid.New())
}

func Parse(s string) (ID, error) {
	id, err := uuid.Parse(s)

	return ID(id), err
}
