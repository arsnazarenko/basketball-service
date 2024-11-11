package apperrors

import "errors"

var (
	ErrPlayerNotFound          = errors.New("player not found")
	ErrInvalidPlayerPageSize   = errors.New("invalid page size for listing player")
	ErrInvalidPlayerPageNumber = errors.New("invalid page number for listing player")
)
