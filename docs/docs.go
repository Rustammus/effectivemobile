// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache helicopter",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/people": {
            "get": {
                "description": "List Peoples by filter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "People API"
                ],
                "summary": "List Peoples by filter",
                "parameters": [
                    {
                        "type": "string",
                        "example": "г. Москва, ул. Ленина, д. 5, кв. 1",
                        "name": "address",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "Иван",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 567890,
                        "name": "passportNumber",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1234,
                        "name": "passportSerie",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "Иванович",
                        "name": "patronymic",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "Иванов",
                        "name": "surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "uuid",
                        "example": "550e8400-e29b-41d4-a716-446655440000",
                        "name": "uuid",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBasePaginated-schemas_ResponsePeople"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    }
                }
            },
            "post": {
                "description": "Create people",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "People API"
                ],
                "summary": "Create people",
                "parameters": [
                    {
                        "description": "People base",
                        "name": "People",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.RequestCreatePeople"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBase-schemas_ResponseUUID"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    }
                }
            }
        },
        "/people/{uuid}": {
            "get": {
                "description": "Find People by uuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "People API"
                ],
                "summary": "Find People by uuid",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "People UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBase-schemas_ResponsePeople"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    }
                }
            },
            "put": {
                "description": "Update people",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "People API"
                ],
                "summary": "Update people",
                "parameters": [
                    {
                        "description": "People base",
                        "name": "UpdatePeople",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.RequestUpdatePeople"
                        }
                    },
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "People UUID",
                        "name": "uuid",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBase-schemas_ResponsePeople"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete people",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "People API"
                ],
                "summary": "Delete people",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "People UUID",
                        "name": "uuid",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBase-schemas_ResponseUUID"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    }
                }
            }
        },
        "/people/{uuid}/start-task": {
            "post": {
                "description": "Create and start task by People uuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "People API"
                ],
                "summary": "Create and start task by People uuid",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "People UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Task name",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBase-schemas_ResponseUUID"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    }
                }
            }
        },
        "/people/{uuid}/tasks": {
            "get": {
                "description": "List all tasks by People uuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "People API"
                ],
                "summary": "List all tasks by People uuid",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "People UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseMulti-schemas_ResponseTask"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    }
                }
            }
        },
        "/task": {
            "get": {
                "description": "List tasks by people",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task API"
                ],
                "summary": "List tasks by people",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "People UUID",
                        "name": "people",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseMulti-schemas_ResponseTask"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    }
                }
            },
            "post": {
                "description": "Create and start task by People uuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task API"
                ],
                "summary": "Create and start task by People uuid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Task name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "People",
                        "name": "people",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBase-schemas_ResponseUUID"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    }
                }
            }
        },
        "/task/{uuid}": {
            "put": {
                "description": "Stop task by uuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task API"
                ],
                "summary": "Stop task by uuid",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Task UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBase-schemas_ResponseTask"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.IResponseBaseErr"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "crud.Pagination": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                }
            }
        },
        "schemas.RequestCreatePeople": {
            "type": "object",
            "properties": {
                "passportNumber": {
                    "type": "string"
                }
            }
        },
        "schemas.RequestUpdatePeople": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "г. Москва, ул. Ленина, д. 5, кв. 1"
                },
                "name": {
                    "type": "string",
                    "example": "Иван"
                },
                "passportNumber": {
                    "type": "integer",
                    "example": 567890
                },
                "passportSerie": {
                    "type": "integer",
                    "example": 1234
                },
                "patronymic": {
                    "type": "string",
                    "example": "Иванович"
                },
                "surname": {
                    "type": "string",
                    "example": "Иванов"
                }
            }
        },
        "schemas.ResponsePeople": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "г. Москва, ул. Ленина, д. 5, кв. 1"
                },
                "created_at": {
                    "type": "string",
                    "example": "2020-01-01T00:00:00Z"
                },
                "name": {
                    "type": "string",
                    "example": "Иван"
                },
                "passportNumber": {
                    "type": "integer",
                    "example": 567890
                },
                "passportSerie": {
                    "type": "integer",
                    "example": 1234
                },
                "patronymic": {
                    "type": "string",
                    "example": "Иванович"
                },
                "surname": {
                    "type": "string",
                    "example": "Иванов"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2020-01-01T00:00:00Z"
                },
                "uuid": {
                    "type": "string",
                    "example": "550e8400-e29b-41d4-a716-446655440000"
                }
            }
        },
        "schemas.ResponseTask": {
            "type": "object",
            "properties": {
                "hours": {
                    "type": "integer",
                    "example": 25
                },
                "isStopped": {
                    "type": "boolean",
                    "example": true
                },
                "minutes": {
                    "type": "integer",
                    "example": 59
                },
                "name": {
                    "type": "string",
                    "example": "do nothing"
                },
                "people_uuid": {
                    "type": "string",
                    "example": "550e8400-e29b-41d4-a716-446655440000"
                },
                "uuid": {
                    "type": "string",
                    "example": "550e8400-e29b-41d4-a716-446655440000"
                }
            }
        },
        "schemas.ResponseUUID": {
            "type": "object",
            "properties": {
                "uuid": {
                    "type": "string",
                    "example": "550e8400-e29b-41d4-a716-446655440000"
                }
            }
        },
        "v1.IResponseBase-schemas_ResponsePeople": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.ResponsePeople"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.IResponseBase-schemas_ResponseTask": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.ResponseTask"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.IResponseBase-schemas_ResponseUUID": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.ResponseUUID"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.IResponseBaseErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.IResponseBaseMulti-schemas_ResponseTask": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.ResponseTask"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.IResponseBasePaginated-schemas_ResponsePeople": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.ResponsePeople"
                    }
                },
                "message": {
                    "type": "string"
                },
                "next_pagination": {
                    "$ref": "#/definitions/crud.Pagination"
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8082",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "UwU",
	Description:      "This is my server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
