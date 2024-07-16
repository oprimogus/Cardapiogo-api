// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Gustavo Ferreira de Jesus",
            "email": "gustavo081900@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/auth/sign-in": {
            "post": {
                "description": "Authenticate a user using email and password and issue a JWT on successful login.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Sign-In with email and password",
                "parameters": [
                    {
                        "description": "SignInParams",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/authentication.SignInParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/object.JWT"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/auth/sign-up": {
            "post": {
                "description": "Sign-Up with local credentials and data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Sign-Up with local credentials and data",
                "parameters": [
                    {
                        "description": "CreateUserParams",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/user.CreateParams"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/store": {
            "put": {
                "description": "Owner can update your stores.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Store"
                ],
                "summary": "Owner can update your stores.",
                "parameters": [
                    {
                        "description": "Params to update a store",
                        "name": "Params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/store.UpdateParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Owner user can create store",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Store"
                ],
                "summary": "Owner can create stores.",
                "parameters": [
                    {
                        "description": "Params to create a store",
                        "name": "Params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/store.CreateParams"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/store/business-hours": {
            "put": {
                "description": "Owner can update business hours of store.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Store"
                ],
                "summary": "Owner can update business hours of store.",
                "parameters": [
                    {
                        "description": "Params to update business hours of store",
                        "name": "Params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/store.StoreBusinessHoursParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Owner can delete business hours of store.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Store"
                ],
                "summary": "Owner can delete business hours of store.",
                "parameters": [
                    {
                        "description": "Params to delete business hours of store",
                        "name": "Params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/store.StoreBusinessHoursParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/store/{id}": {
            "get": {
                "description": "Any user can view a store.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Store"
                ],
                "summary": "Any user can view a store.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Store ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/store.GetStoreByIdOutput"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/user": {
            "put": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "Update user profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user profile",
                "parameters": [
                    {
                        "description": "UpdateProfileParams",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/user.UpdateProfileParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/user/roles": {
            "post": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "Add a new role for user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Add a new role for user",
                "parameters": [
                    {
                        "description": "AddRolesParams",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/user.AddRolesParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/xerrors.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "authentication.SignInParams": {
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
        "entity.BusinessHours": {
            "type": "object",
            "required": [
                "closingTime",
                "openingTime"
            ],
            "properties": {
                "closingTime": {
                    "type": "string"
                },
                "openingTime": {
                    "type": "string"
                },
                "weekDay": {
                    "type": "integer"
                }
            }
        },
        "entity.PaymentMethod": {
            "type": "string",
            "enum": [
                "credit",
                "debit",
                "pix",
                "cash"
            ],
            "x-enum-varnames": [
                "Credit",
                "Debit",
                "Pix",
                "Cash"
            ]
        },
        "entity.ShopType": {
            "type": "string",
            "enum": [
                "restaurant",
                "pharmacy",
                "tobbaco",
                "market",
                "convenience",
                "pub"
            ],
            "x-enum-varnames": [
                "StoreShopRestaurant",
                "StoreShopPharmacy",
                "StoreShopTobbaco",
                "StoreShopMarket",
                "StoreShopConvenience",
                "StoreShopPub"
            ]
        },
        "entity.UserRole": {
            "type": "string",
            "enum": [
                "consumer",
                "owner",
                "employee",
                "delivery_man",
                "admin"
            ],
            "x-enum-varnames": [
                "UserRoleConsumer",
                "UserRoleOwner",
                "UserRoleEmployee",
                "UserRoleDeliveryMan",
                "UserRoleAdmin"
            ]
        },
        "object.Address": {
            "type": "object",
            "required": [
                "addressLine1",
                "addressLine2",
                "city",
                "country",
                "neighborhood",
                "postalCode",
                "state"
            ],
            "properties": {
                "addressLine1": {
                    "type": "string",
                    "maxLength": 40
                },
                "addressLine2": {
                    "type": "string",
                    "maxLength": 20
                },
                "city": {
                    "type": "string",
                    "maxLength": 25
                },
                "country": {
                    "type": "string",
                    "maxLength": 15
                },
                "latitude": {
                    "type": "string"
                },
                "longitude": {
                    "type": "string"
                },
                "neighborhood": {
                    "type": "string",
                    "maxLength": 25
                },
                "postalCode": {
                    "type": "string",
                    "maxLength": 15
                },
                "state": {
                    "type": "string",
                    "maxLength": 15
                }
            }
        },
        "object.JWT": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_in": {
                    "type": "integer"
                },
                "id_token": {
                    "type": "string"
                },
                "not-before-policy": {
                    "type": "integer"
                },
                "refresh_expires_in": {
                    "type": "integer"
                },
                "refresh_token": {
                    "type": "string"
                },
                "scope": {
                    "type": "string"
                },
                "session_state": {
                    "type": "string"
                },
                "token_type": {
                    "type": "string"
                }
            }
        },
        "store.AddressOutput": {
            "type": "object",
            "properties": {
                "addressLine1": {
                    "type": "string"
                },
                "addressLine2": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "neighborhood": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                }
            }
        },
        "store.CreateParams": {
            "type": "object",
            "required": [
                "address",
                "cpfCnpj",
                "name",
                "phone",
                "type"
            ],
            "properties": {
                "address": {
                    "$ref": "#/definitions/object.Address"
                },
                "cpfCnpj": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 25
                },
                "phone": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/entity.ShopType"
                }
            }
        },
        "store.GetStoreByIdOutput": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/store.AddressOutput"
                },
                "businessHours": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.BusinessHours"
                    }
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "paymentMethod": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.PaymentMethod"
                    }
                },
                "phone": {
                    "type": "string"
                },
                "score": {
                    "type": "integer"
                },
                "type": {
                    "$ref": "#/definitions/entity.ShopType"
                }
            }
        },
        "store.StoreBusinessHoursParams": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "businessHours": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.BusinessHours"
                    }
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "store.UpdateParams": {
            "type": "object",
            "required": [
                "address",
                "id",
                "name",
                "phone",
                "type"
            ],
            "properties": {
                "address": {
                    "$ref": "#/definitions/object.Address"
                },
                "businessHours": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.BusinessHours"
                    }
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 25
                },
                "paymentMethod": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.PaymentMethod"
                    }
                },
                "phone": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/entity.ShopType"
                }
            }
        },
        "user.AddRolesParams": {
            "type": "object",
            "required": [
                "role"
            ],
            "properties": {
                "role": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.UserRole"
                    }
                }
            }
        },
        "user.CreateParams": {
            "type": "object",
            "required": [
                "email",
                "password",
                "profile"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "profile": {
                    "$ref": "#/definitions/user.CreateProfileParams"
                }
            }
        },
        "user.CreateProfileParams": {
            "type": "object",
            "required": [
                "document",
                "lastName",
                "name",
                "phone"
            ],
            "properties": {
                "document": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "user.UpdateProfileParams": {
            "type": "object",
            "required": [
                "lastName",
                "name",
                "phone"
            ],
            "properties": {
                "lastName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "xerrors.ErrorResponse": {
            "type": "object",
            "properties": {
                "debug": {},
                "details": {},
                "error": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer Token": {
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
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Cardapiogo API",
	Description:      "Documentação da API de delivery Cardapiogo.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
