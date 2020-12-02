package promconfig

import (
	"context"
	"net/http"
)

var (
	StatusErr = "error"
	StatusSuccess = "success"
	HandlerNameID = "handler"
	StatusNameID = "status"
	GetUserTickets = "GetUserTickets"
	GetUsersSimpleTicket = "GetUsersSimpleTicket"
	GetHallScheduleTickets = "GetHallScheduleTickets"
	BuyTicket = "BuyTicket"
	GetMovieSchedule = "GetMovieSchedule"
	GetSchedule = "GetSchedule"
	GetRecommendedMovieList = "GetRecommendedMovieList"
	GetProfile = "GetProfile"
	UpdateProfile = "UpdateProfile"
	GenerateCSRFToken = "GenerateCSRFToken"
	GetHallStructure = "GetHallStructure"
	GetCinema = "GetCinema"
	GetCinemaList = "GetCinemaList"
	SignOutHandler = "SignOutHandler"
	RegisterHandler = "RegisterHandler"
	AuthHandler = "AuthHandler"
	GetMovieList = "GetMovieList"
	GetMovie = "GetMovie"
	RateMovie = "RateMovie"
	GetActualMovies = "GetActualMovies"

)


func SetRequestMonitoringContext(r *http.Request, handlerName, status string){
	ctx := r.Context()
	ctx = context.WithValue(ctx,HandlerNameID, handlerName)
	ctx = context.WithValue(ctx,StatusNameID, status)
	*r = *r.Clone(ctx)
}