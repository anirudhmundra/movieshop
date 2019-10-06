// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-10-06 16:35:53.426833 +0530 IST m=+0.020773682

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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Get All Movies",
                "tags": [
                    "movie"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Movies"
                        }
                    }
                }
            }
        },
        "/info": {
            "get": {
                "description": "Get Project Info",
                "tags": [
                    "info"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Info"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Info": {
            "type": "object",
            "properties": {
                "version": {
                    "type": "string"
                }
            }
        },
        "model.Movie": {
            "type": "object",
            "properties": {
                "director": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "model.Movies": {
            "type": "object",
            "properties": {
                "movies": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Movie"
                    }
                }
            }
        }
    },
    "tags": [
        {
            "description": "Gives information of the API",
            "name": "info"
        },
        {
            "description": "Gives movies and information",
            "name": "movie"
        }
    ]
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
	Version:     "v1.0",
	Host:        "",
	BasePath:    "/api/movies",
	Schemes:     []string{},
	Title:       "Movie Shop APIs",
	Description: "List of Movie APIs",
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