// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/sign-in": {
            "post": {
                "description": "user sign-in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "SignIn",
                "operationId": "sign-in",
                "parameters": [
                    {
                        "description": "User info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SignInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.getCreationToken"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "user sign-up",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "SignUp",
                "operationId": "sign-up",
                "parameters": [
                    {
                        "description": "User info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SignUpInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.getCreationId"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    }
                }
            }
        },
        "/language/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get languages list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "language"
                ],
                "summary": "Get Languages List",
                "operationId": "get-languages",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.getAllLanguagesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
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
                "description": "create language",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "language"
                ],
                "summary": "Create Language",
                "operationId": "create-language",
                "parameters": [
                    {
                        "description": "language info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Language"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.getCreationId"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    }
                }
            }
        },
        "/language/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get language by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "language"
                ],
                "summary": "Get Language By ID",
                "operationId": "get-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Language"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
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
                "description": "update language by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "language"
                ],
                "summary": "Update Language",
                "operationId": "update-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "language info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UpdateLanguageInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.statusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
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
                "description": "delete language by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "language"
                ],
                "summary": "Delete Language",
                "operationId": "delete-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.statusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/rest.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Language": {
            "type": "object",
            "properties": {
                "date_of_creation": {
                    "type": "integer"
                },
                "developer": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "domain.SignInInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.SignUpInput": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.UpdateLanguageInput": {
            "type": "object",
            "properties": {
                "date_of_creation": {
                    "type": "integer"
                },
                "developer": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "rest.errorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "rest.getAllLanguagesResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Language"
                    }
                }
            }
        },
        "rest.getCreationId": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "rest.getCreationToken": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "rest.statusResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "CRUD API Languages",
	Description:      "API about programming languages",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
