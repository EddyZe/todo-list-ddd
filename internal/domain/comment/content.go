package comment

import (
	"errors"
	"strings"
)

var (
	ErrCommentIsEmpty = errors.New("comment is empty")
)

type Content string

func NewContent(text string) (Content, error) {
	text = strings.TrimSpace(text)

	if text == "" {
		return "", ErrCommentIsEmpty
	}

	return Content(text), nil
}

func (c Content) String() string {
	return string(c)
}
