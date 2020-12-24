package models

import "errors"

var (
	ErrFooNoDBConnection      = errors.New("no database connection")
	ErrFooInternalDBErr       = errors.New("internal service error")
	ErrFooCastErr             = errors.New("cast error")
	ErrFooIncorrectSQLQuery   = errors.New("incorrect SQL Query")
	ErrFooNoAuthorization     = errors.New("no authorization")
	ErrFooPlaceAlreadyBusy    = errors.New("place already busy")
	ErrFooPlaceDoesntExists   = errors.New("place doesnt exists")
	ErrFooNoLoginInfo         = errors.New("no login information")
	ErrFooIncorrectInputInfo  = errors.New("incorrect info")
	ErrFooArgsMismatch        = errors.New("args have nil")
	ErrFooInitFail            = errors.New("init fail")
	ErrFooInternalServerError = errors.New("internal server error")
	ErrFooNoRatingInfo        = errors.New("nothing to show")
	ErrFooUniqueFail          = errors.New("already exists")
	ErrFooIncorrectPath       = errors.New("incorrect path")
)
