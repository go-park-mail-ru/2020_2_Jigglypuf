// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/auth/login/": {
            "post": {
                "description": "login user and get session",
                "consumes": [
                    "application/json"
                ],
                "summary": "login",
                "operationId": "login-user-by-login-data",
                "parameters": [
                    {
                        "description": "Login information",
                        "name": "Login_info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AuthInput"
                        }
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/logout/": {
            "post": {
                "description": "SignOut user",
                "summary": "SignOut",
                "operationId": "SignOut-user-by-register-data",
                "responses": {
                    "200": {},
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/register/": {
            "post": {
                "description": "register user and get session",
                "consumes": [
                    "application/json"
                ],
                "summary": "Register",
                "operationId": "register-user-by-register-data",
                "parameters": [
                    {
                        "description": "Register information",
                        "name": "Register_info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RegistrationInput"
                        }
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            }
        },
        "/api/cinema/": {
            "get": {
                "description": "Get cinema list",
                "summary": "GetCinemaList",
                "operationId": "cinema-list-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Cinema"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            }
        },
        "/api/cinema/{id}/": {
            "get": {
                "description": "Get cinema",
                "summary": "GetCinema",
                "operationId": "cinema-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "cinema id param",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Cinema"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            }
        },
        "/api/csrf/": {
            "get": {
                "description": "Returns movie schedule by ID",
                "summary": "Get CSRF by session",
                "operationId": "csrf-id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/csrf.Response"
                        }
                    },
                    "400": {
                        "description": "Bad body",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "500": {
                        "description": "internal error",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            }
        },
        "/api/hall/{id}/": {
            "get": {
                "description": "Get cinema hall placement structure",
                "summary": "Get hall structure",
                "operationId": "hall-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "hall id param",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CinemaHall"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            }
        },
        "/api/movie/": {
            "get": {
                "description": "Get movie list",
                "summary": "GetMovieList",
                "operationId": "movie-list-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.MovieList"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            }
        },
        "/api/movie/actual/": {
            "get": {
                "description": "Returns movie that in the cinema",
                "summary": "Get movies in cinema",
                "operationId": "movie-in-cinema-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.MovieList"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad body",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "401": {
                        "description": "No authorization",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            }
        },
        "/api/movie/rate/": {
            "post": {
                "description": "Rate movie",
                "consumes": [
                    "application/json"
                ],
                "summary": "RateMovie",
                "operationId": "movie-rate-id",
                "parameters": [
                    {
                        "description": "Login information",
                        "name": "Login_info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RateMovie"
                        }
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {
                        "description": "Bad body",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "401": {
                        "description": "No authorization",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            }
        },
        "/api/movie/{id}/": {
            "get": {
                "description": "Get movie",
                "summary": "GetMovie",
                "operationId": "movie-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "movie id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Movie"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            }
        },
        "/api/profile/": {
            "get": {
                "description": "Get Profile",
                "summary": "GetProfile",
                "operationId": "profile-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cookie information",
                        "name": "Cookie_info",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Profile"
                        }
                    },
                    "400": {
                        "description": "Bad body",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "401": {
                        "description": "No authorization",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Get Profile",
                "summary": "GetProfile",
                "operationId": "profile-update-id",
                "parameters": [
                    {
                        "type": "string",
                        "name": "avatar",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "surname",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {
                        "description": "Bad body",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "401": {
                        "description": "No authorization",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            }
        },
        "/api/schedule/": {
            "get": {
                "description": "Returns movie schedule by getting movie id, cinema id and day(date) in format schedule.TimeStandard",
                "summary": "Get movie schedule",
                "operationId": "movie-schedule-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "movie_id",
                        "name": "movie_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "cinema_id",
                        "name": "cinema_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "date",
                        "name": "date",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Schedule"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad body",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "500": {
                        "description": "internal error",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            }
        },
        "/api/schedule/{id}": {
            "get": {
                "description": "Returns movie schedule by ID",
                "summary": "Get schedule by id",
                "operationId": "schedule-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "schedule id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Schedule"
                        }
                    },
                    "400": {
                        "description": "Bad body",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "500": {
                        "description": "internal error",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            }
        },
        "/api/ticket/": {
            "get": {
                "description": "Get user ticket list",
                "summary": "Get user ticket list",
                "operationId": "get-ticket-list-id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Ticket"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad body",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "401": {
                        "description": "No auth",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal err",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            }
        },
        "/api/ticket/buy/": {
            "post": {
                "description": "Buys ticket by schedule ID and place to authenticated user or by e-mail",
                "consumes": [
                    "application/json"
                ],
                "summary": "Buy ticket",
                "operationId": "buy-ticket-id",
                "parameters": [
                    {
                        "description": "Ticket info",
                        "name": "Ticket_info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TicketInput"
                        }
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {
                        "description": "Bad body",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            }
        },
        "/api/ticket/schedule/{id}/": {
            "get": {
                "description": "Get schedule hall ticket list by id",
                "summary": "Get schedule hall ticket list",
                "operationId": "get-schedule-ticket-list-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "schedule_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.TicketPlace"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad body",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal err",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            }
        },
        "/api/ticket/{id}/": {
            "get": {
                "description": "Get user ticket by id",
                "summary": "Get user ticket",
                "operationId": "get-ticket-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ticket id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Ticket"
                        }
                    },
                    "400": {
                        "description": "Bad body",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "401": {
                        "description": "No auth",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal err",
                        "schema": {
                            "$ref": "#/definitions/models.ServerResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "csrf.Response": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "models.AuthInput": {
            "type": "object",
            "required": [
                "login"
            ],
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.Cinema": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "authorID": {
                    "type": "integer"
                },
                "hallCount": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "pathToAvatar": {
                    "type": "string"
                }
            }
        },
        "models.CinemaHall": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "placeAmount": {
                    "type": "integer"
                },
                "placeConfig": {
                    "$ref": "#/definitions/models.HallConfig"
                }
            }
        },
        "models.HallConfig": {
            "type": "object",
            "properties": {
                "levels": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.HallPlace"
                    }
                }
            }
        },
        "models.HallPlace": {
            "type": "object",
            "properties": {
                "place": {
                    "type": "integer"
                },
                "row": {
                    "type": "integer"
                }
            }
        },
        "models.Movie": {
            "type": "object",
            "properties": {
                "actors": {
                    "type": "string"
                },
                "ageGroup": {
                    "type": "integer"
                },
                "country": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "genre": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "pathToAvatar": {
                    "type": "string"
                },
                "pathToSliderAvatar": {
                    "type": "string"
                },
                "personalRating": {
                    "type": "integer"
                },
                "producer": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "ratingCount": {
                    "type": "integer"
                },
                "releaseYear": {
                    "type": "integer"
                }
            }
        },
        "models.MovieList": {
            "type": "object",
            "properties": {
                "actors": {
                    "type": "string"
                },
                "ageGroup": {
                    "type": "integer"
                },
                "country": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "genre": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "pathToAvatar": {
                    "type": "string"
                },
                "pathToSliderAvatar": {
                    "type": "string"
                },
                "producer": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "ratingCount": {
                    "type": "integer"
                },
                "releaseYear": {
                    "type": "integer"
                }
            }
        },
        "models.Profile": {
            "type": "object",
            "properties": {
                "avatarPath": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                },
                "userCredentials": {
                    "$ref": "#/definitions/models.User"
                }
            }
        },
        "models.RateMovie": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "integer"
                }
            }
        },
        "models.RegistrationInput": {
            "type": "object",
            "required": [
                "login"
            ],
            "properties": {
                "login": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "models.Schedule": {
            "type": "object",
            "properties": {
                "cinemaID": {
                    "type": "integer"
                },
                "cost": {
                    "type": "integer"
                },
                "hallID": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "movieID": {
                    "type": "integer"
                },
                "premierTime": {
                    "type": "string"
                }
            }
        },
        "models.ServerResponse": {
            "type": "object",
            "properties": {
                "response": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "models.Ticket": {
            "type": "object",
            "required": [
                "login"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "login": {
                    "type": "string"
                },
                "placeField": {
                    "$ref": "#/definitions/models.TicketPlace"
                },
                "schedule": {
                    "$ref": "#/definitions/models.Schedule"
                },
                "transactionDate": {
                    "type": "string"
                }
            }
        },
        "models.TicketInput": {
            "type": "object",
            "required": [
                "login"
            ],
            "properties": {
                "login": {
                    "type": "string"
                },
                "placeField": {
                    "$ref": "#/definitions/models.TicketPlace"
                },
                "scheduleID": {
                    "type": "integer"
                }
            }
        },
        "models.TicketPlace": {
            "type": "object",
            "properties": {
                "place": {
                    "type": "integer"
                },
                "row": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "required": [
                "login"
            ],
            "properties": {
                "login": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.5",
	Host:        "https://cinemascope.space",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "CinemaScope Backend API",
	Description: "This is a backend API",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
