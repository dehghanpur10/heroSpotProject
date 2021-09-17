// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Mohammad Dehghanpour",
            "email": "m.dehghanpour10@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v2/reservations": {
            "get": {
                "description": "this endpoint Get the summary of all reservations",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservation"
                ],
                "summary": "Get the summary of all reservations",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Reservation"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/lib.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "this endpoint creates a new reservation for vehicle",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservation"
                ],
                "summary": "create a new reservation for vehicle",
                "parameters": [
                    {
                        "description": "vehicle info",
                        "name": "reservation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.InputReservation"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Reservation"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/lib.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/lib.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/reservations/{reservation_id}/update": {
            "get": {
                "description": "this endpoint will check  possibility for update time",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservation"
                ],
                "summary": "checking  possibility  for update time",
                "parameters": [
                    {
                        "type": "string",
                        "description": "reservation ID",
                        "name": "reservation_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Reservation"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/lib.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/lib.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/lib.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "this endpoint will update reservation time",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservation"
                ],
                "summary": "update reservation time",
                "parameters": [
                    {
                        "type": "string",
                        "description": "reservation ID",
                        "name": "reservation_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "reservation info",
                        "name": "reservation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.InputReservation"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Reservation"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/lib.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/lib.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/lib.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/search": {
            "get": {
                "description": "if user enter lon and lat query this endpoint will search facility based on lon and lat facility if user don't enter, this endpoint will send all facility",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search"
                ],
                "summary": "search facility",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Latitude",
                        "name": "lat",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "longitude",
                        "name": "lon",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Facility"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/lib.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/lib.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/vehicles": {
            "post": {
                "description": "this endpoint creates a new vehicle for user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vehicle"
                ],
                "summary": "create a new vehicle for user",
                "parameters": [
                    {
                        "description": "vehicle info",
                        "name": "vehicle",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Vehicle"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Vehicle"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/lib.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/lib.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "lib.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.Facility": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "facility_id": {
                    "type": "string"
                },
                "latitude": {
                    "type": "integer"
                },
                "longitude": {
                    "type": "integer"
                }
            }
        },
        "models.InputReservation": {
            "type": "object",
            "required": [
                "facility_id",
                "parked_vehicle_id",
                "reservation_id",
                "update_possible"
            ],
            "properties": {
                "facility_id": {
                    "type": "string",
                    "example": "1"
                },
                "parked_vehicle_id": {
                    "type": "string",
                    "example": "1"
                },
                "quote": {
                    "$ref": "#/definitions/models.Quote"
                },
                "reservation_id": {
                    "type": "string",
                    "example": "1"
                },
                "update_possible": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.Quote": {
            "type": "object",
            "required": [
                "ends",
                "starts"
            ],
            "properties": {
                "ends": {
                    "type": "string",
                    "example": "2019-08-19T13:49:37.000Z"
                },
                "starts": {
                    "type": "string",
                    "example": "2019-08-19T13:49:37.000Z"
                }
            }
        },
        "models.Reservation": {
            "type": "object",
            "properties": {
                "facility": {
                    "$ref": "#/definitions/models.Facility"
                },
                "parked_vehicle": {
                    "$ref": "#/definitions/models.Vehicle"
                },
                "quote": {
                    "$ref": "#/definitions/models.Quote"
                },
                "reservation_id": {
                    "type": "string"
                },
                "update_possible": {
                    "type": "boolean"
                }
            }
        },
        "models.Vehicle": {
            "type": "object",
            "required": [
                "vehicle_description",
                "vehicle_id"
            ],
            "properties": {
                "vehicle_description": {
                    "$ref": "#/definitions/models.VehicleDescription"
                },
                "vehicle_id": {
                    "type": "string",
                    "example": "1"
                }
            }
        },
        "models.VehicleDescription": {
            "type": "object",
            "required": [
                "model",
                "name",
                "year"
            ],
            "properties": {
                "model": {
                    "type": "string",
                    "example": "s300"
                },
                "name": {
                    "type": "string",
                    "example": "benz"
                },
                "year": {
                    "type": "string",
                    "example": "2021"
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
	Version:     "2.0",
	Host:        "tvwvnaqfy9.execute-api.us-west-2.amazonaws.com/api",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Spot Hero",
	Description: "Implement spot hero",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
