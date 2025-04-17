package docs

import (
	"github.com/swaggo/swag"
)

// @title Movies API
// @version 1.0
// @description API для управления фильмами
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func init() {
	swag.Register(swag.Name, &s{})
}

type s struct{}

func (s *s) ReadDoc() string {
	doc := `{
    "swagger": "2.0",
    "info": {
        "description": "API для управления фильмами",
        "version": "1.0",
        "title": "Movies API",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        }
    },
    "host": "localhost:8080",
    "basePath": "/api",
    "schemes": [
        "http"
    ],
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "paths": {
        "/movies": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Получить список всех фильмов",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Movie"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Создать новый фильм",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "movie",
                        "description": "Данные фильма",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Movie"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/Movie"
                        }
                    }
                }
            }
        },
        "/movies/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Получить фильм по ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID фильма",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Movie"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Обновить фильм",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID фильма",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "in": "body",
                        "name": "movie",
                        "description": "Обновленные данные фильма",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Movie"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Movie"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Удалить фильм",
                "tags": [
                    "movies"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID фильма",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "Movie": {
            "type": "object",
            "required": [
                "title",
                "director",
                "year"
            ],
            "properties": {
                "id": {
                    "type": "integer",
                    "format": "int64"
                },
                "title": {
                    "type": "string"
                },
                "director": {
                    "type": "string"
                },
                "year": {
                    "type": "integer",
                    "format": "int32"
                },
                "plot": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string",
                    "format": "date-time"
                },
                "updated_at": {
                    "type": "string",
                    "format": "date-time"
                }
            }
        }
    }
}`
	return doc
}
