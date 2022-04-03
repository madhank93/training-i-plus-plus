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
        "/appointment": {
            "post": {
                "description": "create Appointment",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Appointment"
                ],
                "summary": "create Appointment",
                "parameters": [
                    {
                        "description": "appointment info",
                        "name": "appointment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_entities.AppointmentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/appointment/user/{userid}": {
            "get": {
                "description": "Get Appointment By  User",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Appointment"
                ],
                "summary": "Get Appointment By User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "userid",
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
                                "$ref": "#/definitions/controller.UserAppResponse"
                            }
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/doctor": {
            "get": {
                "description": "fetch all doctor",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Doctor"
                ],
                "summary": "fetch all  doctor",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Doctor"
                            }
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "post": {
                "description": "create doctor",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Doctor"
                ],
                "summary": "create doctor",
                "parameters": [
                    {
                        "description": "doctor info",
                        "name": "doctor",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.DoctorRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/doctor/{_id}": {
            "patch": {
                "description": "Update doctor",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Doctor"
                ],
                "summary": "Update doctor",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "doctor info",
                        "name": "doctor",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.DoctorRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/user/{_id}": {
            "delete": {
                "tags": [
                    "Doctor"
                ],
                "summary": "Delete doctor by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "api_entities.AppointmentRequest": {
            "description": "Appointment Request",
            "type": "object",
            "properties": {
                "doctor": {
                    "type": "string"
                },
                "from": {
                    "type": "string",
                    "example": "02 Jan 22 15:00 IST"
                },
                "patient": {
                    "$ref": "#/definitions/api_entities.User"
                },
                "to": {
                    "type": "string",
                    "example": "02 Jan 22 16:00 IST"
                }
            }
        },
        "api_entities.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controller.DoctorRequest": {
            "description": "Doctor Request",
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "qualification": {
                    "type": "string"
                }
            }
        },
        "controller.UserAppResponse": {
            "type": "object",
            "properties": {
                "doctor": {
                    "type": "string"
                },
                "from": {
                    "type": "string",
                    "example": "02 Jan 22 15:00 IST"
                },
                "to": {
                    "type": "string",
                    "example": "02 Jan 22 16:00 IST"
                }
            }
        },
        "models.Appointment": {
            "description": "Appointment",
            "type": "object",
            "required": [
                "from",
                "patient"
            ],
            "properties": {
                "from": {
                    "type": "integer"
                },
                "patient": {
                    "type": "string"
                },
                "to": {
                    "type": "integer"
                }
            }
        },
        "models.Doctor": {
            "description": "Doctor",
            "type": "object",
            "required": [
                "name",
                "qualification"
            ],
            "properties": {
                "_id": {
                    "type": "string"
                },
                "appointments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Appointment"
                    }
                },
                "name": {
                    "type": "string"
                },
                "qualification": {
                    "type": "string"
                },
                "updated_on": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7451",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Doctor Service",
	Description:      "This is doctor crud service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
