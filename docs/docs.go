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
        "/auth/login/": {
            "post": {
                "description": "login user and get cookie",
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
        "/auth/logout/": {
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
        "/auth/register/": {
            "post": {
                "description": "register user and get cookie",
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
        "/cinema/": {
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
        "/cinema/{id}/": {
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
        "/movie/": {
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
                                "$ref": "#/definitions/models.Movie"
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
        "/movie/rate/": {
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
        "/movie/{id}/": {
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
        "/profile/": {
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
        }
    },
    "definitions": {
        "models.AuthInput": {
            "type": "object",
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
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Movie": {
            "type": "object",
            "properties": {
                "description": {
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
                "rating": {
                    "type": "number"
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
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
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
