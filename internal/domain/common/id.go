package common

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidID = errors.New("invalid TaskID")
)

type ID uuid.UUID

func GenerateID() ID {
	return ID(uuid.New())
}

func NewID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return ID{}, ErrInvalidID
	}

	return ID(id), nil
}

func (o ID) String() string {
	return uuid.UUID(o).String()
}
