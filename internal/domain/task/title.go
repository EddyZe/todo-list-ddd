package task

import (
	"errors"
	"strings"
)

var (
	ErrInvalidTitle = errors.New("invalid title")
)

type Title string

func NewTitle(title string) (Title, error) {
	title = strings.TrimSpace(title)

	if title == "" {
		return "", ErrInvalidTitle
	}

	return Title(title), nil
}

func (t Title) String() string {
	return string(t)
}
