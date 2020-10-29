package models

import "errors"

var (
	ErrFooNoDBConnection = errors.New("no database connection")
)
