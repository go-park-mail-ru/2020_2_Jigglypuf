package models

import "errors"

var (
	ErrFooNoDBConnection = errors.New("no database connection")
	ErrFooInternalDBErr  = errors.New("internal service error")
	ErrFooCastErr		 = errors.New("cast error")
)
