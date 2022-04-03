// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/github.com/dhi13man/healthcare-app.v1+json"
  ],
  "produces": [
    "application/github.com/dhi13man/healthcare-app.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "A Simple Healthcare app built to illustrate the concepts learnt during Swigggy i++ Training.",
    "title": "Healthcare App",
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "base"
        ],
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the two microservices",
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/bookkeeping/": {
      "get": {
        "tags": [
          "bookkeeping-service"
        ],
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "The bookkeeping-service endpoint",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/medicine"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/users/": {
      "get": {
        "tags": [
          "users-service"
        ],
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "The users-service endpoint",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/client"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "client": {
      "type": "object",
      "required": [
        "email"
      ],
      "properties": {
        "email": {
          "type": "string",
          "format": "email",
          "minLength": 2,
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        },
        "subscriptions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/user_subscription"
          }
        }
      }
    },
    "disease": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "name": {
          "type": "string",
          "minLength": 1
        },
        "next_dose": {
          "type": "string",
          "format": "string"
        },
        "rate": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "medicine": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "diseases": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "user_subscription": {
      "type": "object",
      "required": [
        "medicine",
        "rate"
      ],
      "properties": {
        "medicine": {
          "type": "string",
          "minLength": 1
        },
        "next_dose": {
          "type": "string",
          "format": "string"
        },
        "rate": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/github.com/dhi13man/healthcare-app.v1+json"
  ],
  "produces": [
    "application/github.com/dhi13man/healthcare-app.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "A Simple Healthcare app built to illustrate the concepts learnt during Swigggy i++ Training.",
    "title": "Healthcare App",
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "base"
        ],
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the two microservices",
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/bookkeeping/": {
      "get": {
        "tags": [
          "bookkeeping-service"
        ],
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "The bookkeeping-service endpoint",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/medicine"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/users/": {
      "get": {
        "tags": [
          "users-service"
        ],
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "The users-service endpoint",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/client"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "client": {
      "type": "object",
      "required": [
        "email"
      ],
      "properties": {
        "email": {
          "type": "string",
          "format": "email",
          "minLength": 2,
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        },
        "subscriptions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/user_subscription"
          }
        }
      }
    },
    "disease": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "name": {
          "type": "string",
          "minLength": 1
        },
        "next_dose": {
          "type": "string",
          "format": "string"
        },
        "rate": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "medicine": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "diseases": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "user_subscription": {
      "type": "object",
      "required": [
        "medicine",
        "rate"
      ],
      "properties": {
        "medicine": {
          "type": "string",
          "minLength": 1
        },
        "next_dose": {
          "type": "string",
          "format": "string"
        },
        "rate": {
          "type": "string"
        }
      }
    }
  }
}`))
}
