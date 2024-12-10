package utils

import "errors"

var (
	ErrInvalidAuthorId    = errors.New("invalid author id")
	ErrDuplicateBookTitle = errors.New("duplicate book title")
)

func IsUnprocessableEntityErrors(err error) bool {
	if errors.Is(err, ErrInvalidAuthorId) {
		return true
	}
	if errors.Is(err, ErrDuplicateBookTitle) {
		return true
	}

	return false
}
