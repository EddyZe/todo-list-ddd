package task

import (
	"errors"
	"strings"
)

var (
	ErrInvalidDescription = errors.New("invalid description")
)

type Description string

func NewDescription(d string) (Description, error) {
	d = strings.TrimSpace(d)

	if d == "" {
		return Description(""), ErrInvalidDescription
	}

	return Description(d), nil
}

func (d Description) String() string {
	return string(d)
}
