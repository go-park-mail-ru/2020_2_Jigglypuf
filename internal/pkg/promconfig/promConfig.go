package promconfig

import (
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


func SetRequestMonitoringContext(r http.ResponseWriter, handlerName string, status *string){
	r.Header().Set(HandlerNameID, handlerName)
	r.Header().Set(StatusNameID, *status)
}