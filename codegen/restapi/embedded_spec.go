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
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Serves Hearthstone cards",
    "title": "Hearthstone Card Service API",
    "contact": {
      "name": "William Winkler",
      "email": "williambwinkler@gmail.com"
    },
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/cards": {
      "get": {
        "description": "Returns cards as they are stored",
        "tags": [
          "cards"
        ],
        "parameters": [
          {
            "$ref": "#/parameters/cardNameParam"
          },
          {
            "$ref": "#/parameters/manaCostParam"
          },
          {
            "$ref": "#/parameters/healthParam"
          },
          {
            "$ref": "#/parameters/attackParam"
          },
          {
            "$ref": "#/parameters/classParam"
          },
          {
            "$ref": "#/parameters/rarityParam"
          },
          {
            "$ref": "#/parameters/typeParam"
          },
          {
            "$ref": "#/parameters/keywordsParam"
          },
          {
            "$ref": "#/parameters/setParam"
          },
          {
            "$ref": "#/parameters/pageParam"
          },
          {
            "$ref": "#/parameters/limitParam"
          }
        ],
        "responses": {
          "200": {
            "description": "Returns the cards based on query. If there is no query, cards will be returned based on their manaCost in ascending order.",
            "schema": {
              "$ref": "#/definitions/cards"
            }
          },
          "500": {
            "description": "Something went wrong internally",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/classes": {
      "get": {
        "description": "Serves the different classes cards can have. Fx \"Warlock\" or \"Neutral\"",
        "tags": [
          "classes"
        ],
        "responses": {
          "200": {
            "description": "Returns all classes",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/classes"
              }
            }
          },
          "500": {
            "description": "Something went wrong internally",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/keywords": {
      "get": {
        "description": "Serves the different keywords cards can have. Fx \"Taunt\" or \"Quest\"",
        "tags": [
          "keywords"
        ],
        "responses": {
          "200": {
            "description": "Returns all keywords",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/keywords"
              }
            }
          },
          "500": {
            "description": "Something went wrong internally",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/rarities": {
      "get": {
        "description": "Serves the different rarities a card can have. Fx \"Common\" or \"Legendary\"",
        "tags": [
          "rarities"
        ],
        "responses": {
          "200": {
            "description": "Returns all rarities",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/rarities"
              }
            }
          },
          "500": {
            "description": "Something went wrong internally",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/richcards": {
      "get": {
        "description": "Rich cards contains names instead of ids of fx CardType \"Minion\", \"Spell\", \"Secret\" etc",
        "tags": [
          "cards"
        ],
        "parameters": [
          {
            "$ref": "#/parameters/cardNameParam"
          },
          {
            "$ref": "#/parameters/manaCostParam"
          },
          {
            "$ref": "#/parameters/healthParam"
          },
          {
            "$ref": "#/parameters/attackParam"
          },
          {
            "$ref": "#/parameters/classParam"
          },
          {
            "$ref": "#/parameters/rarityParam"
          },
          {
            "$ref": "#/parameters/typeParam"
          },
          {
            "$ref": "#/parameters/keywordsParam"
          },
          {
            "$ref": "#/parameters/setParam"
          },
          {
            "$ref": "#/parameters/pageParam"
          },
          {
            "$ref": "#/parameters/limitParam"
          }
        ],
        "responses": {
          "200": {
            "description": "Returns the cards based on query. If there is no query, cards will be returned based on their manaCost in ascending order.",
            "schema": {
              "$ref": "#/definitions/richCards"
            }
          },
          "500": {
            "description": "Something went wrong internally",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/sets": {
      "get": {
        "description": "Cards can belong to different sets or expansions. This serves all sets and their info.",
        "tags": [
          "sets"
        ],
        "responses": {
          "200": {
            "description": "Returns all sets",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/sets"
              }
            }
          },
          "500": {
            "description": "Something went wrong internally",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/types": {
      "get": {
        "description": "Serves the different types cards can be. Fx \"Minion\" or \"Spell\"",
        "tags": [
          "types"
        ],
        "responses": {
          "200": {
            "description": "Returns all types",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/types"
              }
            }
          },
          "500": {
            "description": "Something went wrong internally",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/update": {
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "description": "Checks for updates to Hearthstone, if there are any, they are applied to the database",
        "tags": [
          "update"
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "401": {
            "description": "Unauthorized request"
          }
        }
      }
    }
  },
  "definitions": {
    "card": {
      "type": "object",
      "properties": {
        "artistName": {
          "type": "string",
          "x-omitempty": false
        },
        "attack": {
          "type": "integer",
          "x-omitempty": false
        },
        "cardSetId": {
          "type": "integer",
          "x-omitempty": false
        },
        "cardTypeId": {
          "type": "integer",
          "x-omitempty": false
        },
        "classId": {
          "type": "integer",
          "x-omitempty": false
        },
        "collectible": {
          "type": "integer",
          "x-omitempty": false
        },
        "duals": {
          "$ref": "#/definitions/duals"
        },
        "flavorText": {
          "type": "string"
        },
        "health": {
          "type": "integer",
          "x-omitempty": false
        },
        "id": {
          "description": "This is the ID from blizzards API",
          "type": "integer"
        },
        "image": {
          "description": "Links to a png-image of the card",
          "type": "string"
        },
        "imageGold": {
          "description": "Links to a png-image of the golden version of the card",
          "type": "string"
        },
        "keywordIds": {
          "type": "array",
          "items": {
            "type": "integer"
          }
        },
        "manaCost": {
          "type": "integer",
          "x-omitempty": false
        },
        "name": {
          "type": "string"
        },
        "parentId": {
          "type": "integer",
          "x-omitempty": false
        },
        "rarityId": {
          "type": "integer",
          "x-omitempty": false
        },
        "text": {
          "type": "string"
        }
      }
    },
    "cards": {
      "type": "object",
      "properties": {
        "cardCount": {
          "type": "integer",
          "x-omitempty": false
        },
        "cards": {
          "description": "if there a no cards, the array is null",
          "type": "array",
          "items": {
            "x-omitempty": false,
            "$ref": "#/definitions/card"
          }
        },
        "page": {
          "type": "integer",
          "x-omitempty": false
        },
        "pageCount": {
          "type": "integer",
          "x-omitempty": false
        }
      }
    },
    "classes": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "duals": {
      "type": "object",
      "properties": {
        "constructed": {
          "type": "boolean"
        },
        "relevant": {
          "type": "boolean"
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
          "description": "HTTPS reponse 400+",
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "info": {
      "type": "object",
      "properties": {
        "amountOfCards": {
          "type": "integer",
          "x-omitempty": false
        },
        "lastUpdate": {
          "description": "formatted as RFC 3339",
          "type": "string",
          "format": "date-time"
        },
        "systemStartTime": {
          "description": "formatted as RFC 3339",
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "keywords": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "text": {
          "type": "string"
        }
      }
    },
    "rarities": {
      "type": "object",
      "properties": {
        "craftingcost": {
          "type": "array",
          "items": {
            "type": "integer"
          }
        },
        "dustvalue": {
          "type": "array",
          "items": {
            "type": "integer"
          }
        },
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "richCard": {
      "type": "object",
      "properties": {
        "artistName": {
          "type": "string",
          "x-omitempty": false
        },
        "attack": {
          "type": "integer",
          "x-omitempty": false
        },
        "cardSet": {
          "type": "string",
          "x-omitempty": false
        },
        "cardType": {
          "type": "string",
          "x-omitempty": false
        },
        "class": {
          "type": "string",
          "x-omitempty": false
        },
        "collectible": {
          "type": "integer",
          "x-omitempty": false
        },
        "duals": {
          "$ref": "#/definitions/duals"
        },
        "flavorText": {
          "type": "string"
        },
        "health": {
          "type": "integer",
          "x-omitempty": false
        },
        "id": {
          "description": "This is the ID from blizzards API",
          "type": "integer"
        },
        "image": {
          "description": "Links to a png-image of the card",
          "type": "string"
        },
        "imageGold": {
          "description": "Links to a png-image of the golden version of the card",
          "type": "string"
        },
        "keywords": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "manaCost": {
          "type": "integer",
          "x-omitempty": false
        },
        "name": {
          "type": "string"
        },
        "parentId": {
          "type": "integer",
          "x-omitempty": false
        },
        "rarity": {
          "type": "string",
          "x-omitempty": false
        },
        "text": {
          "type": "string"
        }
      }
    },
    "richCards": {
      "type": "object",
      "properties": {
        "cardCount": {
          "type": "integer"
        },
        "cards": {
          "description": "if there a no cards, the array is null",
          "type": "array",
          "items": {
            "$ref": "#/definitions/richCard"
          },
          "example": null
        },
        "page": {
          "type": "integer"
        },
        "pageCount": {
          "type": "integer"
        }
      }
    },
    "sets": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "types": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        }
      }
    }
  },
  "parameters": {
    "attackParam": {
      "type": "integer",
      "name": "attack",
      "in": "query"
    },
    "cardNameParam": {
      "minLength": 1,
      "type": "string",
      "name": "name",
      "in": "query"
    },
    "classParam": {
      "maximum": 14,
      "minimum": 1,
      "type": "integer",
      "name": "class",
      "in": "query"
    },
    "healthParam": {
      "type": "integer",
      "name": "health",
      "in": "query"
    },
    "keywordsParam": {
      "type": "array",
      "items": {
        "type": "integer"
      },
      "name": "keywords",
      "in": "query"
    },
    "limitParam": {
      "maximum": 100,
      "minimum": 1,
      "type": "integer",
      "default": 20,
      "name": "limit",
      "in": "query"
    },
    "manaCostParam": {
      "maximum": 99,
      "type": "integer",
      "name": "manaCost",
      "in": "query"
    },
    "pageParam": {
      "minimum": 1,
      "type": "integer",
      "default": 1,
      "name": "page",
      "in": "query"
    },
    "rarityParam": {
      "maximum": 5,
      "minimum": 1,
      "type": "integer",
      "name": "rarity",
      "in": "query"
    },
    "setParam": {
      "type": "integer",
      "name": "set",
      "in": "query"
    },
    "typeParam": {
      "type": "integer",
      "name": "type",
      "in": "query"
    }
  },
  "securityDefinitions": {
    "basicAuth": {
      "type": "basic"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Serves Hearthstone cards",
    "title": "Hearthstone Card Service API",
    "contact": {
      "name": "William Winkler",
      "email": "williambwinkler@gmail.com"
    },
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/cards": {
      "get": {
        "description": "Returns cards as they are stored",
        "tags": [
          "cards"
        ],
        "parameters": [
          {
            "minLength": 1,
            "type": "string",
            "name": "name",
            "in": "query"
          },
          {
            "maximum": 99,
            "minimum": 0,
            "type": "integer",
            "name": "manaCost",
            "in": "query"
          },
          {
            "minimum": 0,
            "type": "integer",
            "name": "health",
            "in": "query"
          },
          {
            "minimum": 0,
            "type": "integer",
            "name": "attack",
            "in": "query"
          },
          {
            "maximum": 14,
            "minimum": 1,
            "type": "integer",
            "name": "class",
            "in": "query"
          },
          {
            "maximum": 5,
            "minimum": 1,
            "type": "integer",
            "name": "rarity",
            "in": "query"
          },
          {
            "type": "integer",
            "name": "type",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "integer"
            },
            "name": "keywords",
            "in": "query"
          },
          {
            "type": "integer",
            "name": "set",
            "in": "query"
          },
          {
            "minimum": 1,
            "type": "integer",
            "default": 1,
            "name": "page",
            "in": "query"
          },
          {
            "maximum": 100,
            "minimum": 1,
            "type": "integer",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Returns the cards based on query. If there is no query, cards will be returned based on their manaCost in ascending order.",
            "schema": {
              "$ref": "#/definitions/cards"
            }
          },
          "500": {
            "description": "Something went wrong internally",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/classes": {
      "get": {
        "description": "Serves the different classes cards can have. Fx \"Warlock\" or \"Neutral\"",
        "tags": [
          "classes"
        ],
        "responses": {
          "200": {
            "description": "Returns all classes",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/classes"
              }
            }
          },
          "500": {
            "description": "Something went wrong internally",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/keywords": {
      "get": {
        "description": "Serves the different keywords cards can have. Fx \"Taunt\" or \"Quest\"",
        "tags": [
          "keywords"
        ],
        "responses": {
          "200": {
            "description": "Returns all keywords",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/keywords"
              }
            }
          },
          "500": {
            "description": "Something went wrong internally",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/rarities": {
      "get": {
        "description": "Serves the different rarities a card can have. Fx \"Common\" or \"Legendary\"",
        "tags": [
          "rarities"
        ],
        "responses": {
          "200": {
            "description": "Returns all rarities",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/rarities"
              }
            }
          },
          "500": {
            "description": "Something went wrong internally",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/richcards": {
      "get": {
        "description": "Rich cards contains names instead of ids of fx CardType \"Minion\", \"Spell\", \"Secret\" etc",
        "tags": [
          "cards"
        ],
        "parameters": [
          {
            "minLength": 1,
            "type": "string",
            "name": "name",
            "in": "query"
          },
          {
            "maximum": 99,
            "minimum": 0,
            "type": "integer",
            "name": "manaCost",
            "in": "query"
          },
          {
            "minimum": 0,
            "type": "integer",
            "name": "health",
            "in": "query"
          },
          {
            "minimum": 0,
            "type": "integer",
            "name": "attack",
            "in": "query"
          },
          {
            "maximum": 14,
            "minimum": 1,
            "type": "integer",
            "name": "class",
            "in": "query"
          },
          {
            "maximum": 5,
            "minimum": 1,
            "type": "integer",
            "name": "rarity",
            "in": "query"
          },
          {
            "type": "integer",
            "name": "type",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "integer"
            },
            "name": "keywords",
            "in": "query"
          },
          {
            "type": "integer",
            "name": "set",
            "in": "query"
          },
          {
            "minimum": 1,
            "type": "integer",
            "default": 1,
            "name": "page",
            "in": "query"
          },
          {
            "maximum": 100,
            "minimum": 1,
            "type": "integer",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Returns the cards based on query. If there is no query, cards will be returned based on their manaCost in ascending order.",
            "schema": {
              "$ref": "#/definitions/richCards"
            }
          },
          "500": {
            "description": "Something went wrong internally",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/sets": {
      "get": {
        "description": "Cards can belong to different sets or expansions. This serves all sets and their info.",
        "tags": [
          "sets"
        ],
        "responses": {
          "200": {
            "description": "Returns all sets",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/sets"
              }
            }
          },
          "500": {
            "description": "Something went wrong internally",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/types": {
      "get": {
        "description": "Serves the different types cards can be. Fx \"Minion\" or \"Spell\"",
        "tags": [
          "types"
        ],
        "responses": {
          "200": {
            "description": "Returns all types",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/types"
              }
            }
          },
          "500": {
            "description": "Something went wrong internally",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/update": {
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "description": "Checks for updates to Hearthstone, if there are any, they are applied to the database",
        "tags": [
          "update"
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "401": {
            "description": "Unauthorized request"
          }
        }
      }
    }
  },
  "definitions": {
    "card": {
      "type": "object",
      "properties": {
        "artistName": {
          "type": "string",
          "x-omitempty": false
        },
        "attack": {
          "type": "integer",
          "x-omitempty": false
        },
        "cardSetId": {
          "type": "integer",
          "x-omitempty": false
        },
        "cardTypeId": {
          "type": "integer",
          "x-omitempty": false
        },
        "classId": {
          "type": "integer",
          "x-omitempty": false
        },
        "collectible": {
          "type": "integer",
          "x-omitempty": false
        },
        "duals": {
          "$ref": "#/definitions/duals"
        },
        "flavorText": {
          "type": "string"
        },
        "health": {
          "type": "integer",
          "x-omitempty": false
        },
        "id": {
          "description": "This is the ID from blizzards API",
          "type": "integer"
        },
        "image": {
          "description": "Links to a png-image of the card",
          "type": "string"
        },
        "imageGold": {
          "description": "Links to a png-image of the golden version of the card",
          "type": "string"
        },
        "keywordIds": {
          "type": "array",
          "items": {
            "type": "integer"
          }
        },
        "manaCost": {
          "type": "integer",
          "x-omitempty": false
        },
        "name": {
          "type": "string"
        },
        "parentId": {
          "type": "integer",
          "x-omitempty": false
        },
        "rarityId": {
          "type": "integer",
          "x-omitempty": false
        },
        "text": {
          "type": "string"
        }
      }
    },
    "cards": {
      "type": "object",
      "properties": {
        "cardCount": {
          "type": "integer",
          "x-omitempty": false
        },
        "cards": {
          "description": "if there a no cards, the array is null",
          "type": "array",
          "items": {
            "x-omitempty": false,
            "$ref": "#/definitions/card"
          }
        },
        "page": {
          "type": "integer",
          "x-omitempty": false
        },
        "pageCount": {
          "type": "integer",
          "x-omitempty": false
        }
      }
    },
    "classes": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "duals": {
      "type": "object",
      "properties": {
        "constructed": {
          "type": "boolean"
        },
        "relevant": {
          "type": "boolean"
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
          "description": "HTTPS reponse 400+",
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "info": {
      "type": "object",
      "properties": {
        "amountOfCards": {
          "type": "integer",
          "x-omitempty": false
        },
        "lastUpdate": {
          "description": "formatted as RFC 3339",
          "type": "string",
          "format": "date-time"
        },
        "systemStartTime": {
          "description": "formatted as RFC 3339",
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "keywords": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "text": {
          "type": "string"
        }
      }
    },
    "rarities": {
      "type": "object",
      "properties": {
        "craftingcost": {
          "type": "array",
          "items": {
            "type": "integer"
          }
        },
        "dustvalue": {
          "type": "array",
          "items": {
            "type": "integer"
          }
        },
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "richCard": {
      "type": "object",
      "properties": {
        "artistName": {
          "type": "string",
          "x-omitempty": false
        },
        "attack": {
          "type": "integer",
          "x-omitempty": false
        },
        "cardSet": {
          "type": "string",
          "x-omitempty": false
        },
        "cardType": {
          "type": "string",
          "x-omitempty": false
        },
        "class": {
          "type": "string",
          "x-omitempty": false
        },
        "collectible": {
          "type": "integer",
          "x-omitempty": false
        },
        "duals": {
          "$ref": "#/definitions/duals"
        },
        "flavorText": {
          "type": "string"
        },
        "health": {
          "type": "integer",
          "x-omitempty": false
        },
        "id": {
          "description": "This is the ID from blizzards API",
          "type": "integer"
        },
        "image": {
          "description": "Links to a png-image of the card",
          "type": "string"
        },
        "imageGold": {
          "description": "Links to a png-image of the golden version of the card",
          "type": "string"
        },
        "keywords": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "manaCost": {
          "type": "integer",
          "x-omitempty": false
        },
        "name": {
          "type": "string"
        },
        "parentId": {
          "type": "integer",
          "x-omitempty": false
        },
        "rarity": {
          "type": "string",
          "x-omitempty": false
        },
        "text": {
          "type": "string"
        }
      }
    },
    "richCards": {
      "type": "object",
      "properties": {
        "cardCount": {
          "type": "integer"
        },
        "cards": {
          "description": "if there a no cards, the array is null",
          "type": "array",
          "items": {
            "$ref": "#/definitions/richCard"
          },
          "example": []
        },
        "page": {
          "type": "integer"
        },
        "pageCount": {
          "type": "integer"
        }
      }
    },
    "sets": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "types": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        }
      }
    }
  },
  "parameters": {
    "attackParam": {
      "minimum": 0,
      "type": "integer",
      "name": "attack",
      "in": "query"
    },
    "cardNameParam": {
      "minLength": 1,
      "type": "string",
      "name": "name",
      "in": "query"
    },
    "classParam": {
      "maximum": 14,
      "minimum": 1,
      "type": "integer",
      "name": "class",
      "in": "query"
    },
    "healthParam": {
      "minimum": 0,
      "type": "integer",
      "name": "health",
      "in": "query"
    },
    "keywordsParam": {
      "type": "array",
      "items": {
        "type": "integer"
      },
      "name": "keywords",
      "in": "query"
    },
    "limitParam": {
      "maximum": 100,
      "minimum": 1,
      "type": "integer",
      "default": 20,
      "name": "limit",
      "in": "query"
    },
    "manaCostParam": {
      "maximum": 99,
      "minimum": 0,
      "type": "integer",
      "name": "manaCost",
      "in": "query"
    },
    "pageParam": {
      "minimum": 1,
      "type": "integer",
      "default": 1,
      "name": "page",
      "in": "query"
    },
    "rarityParam": {
      "maximum": 5,
      "minimum": 1,
      "type": "integer",
      "name": "rarity",
      "in": "query"
    },
    "setParam": {
      "type": "integer",
      "name": "set",
      "in": "query"
    },
    "typeParam": {
      "type": "integer",
      "name": "type",
      "in": "query"
    }
  },
  "securityDefinitions": {
    "basicAuth": {
      "type": "basic"
    }
  }
}`))
}
