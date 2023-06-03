// Code generated by swaggo/swag. DO NOT EDIT.

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
            "url": "lifthus531@gmail.com",
            "email": "lifthus531@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/program/weekly": {
            "post": {
                "tags": [
                    "program"
                ],
                "summary": "gets CreateWeeklyProgramDto from body and returns created program",
                "parameters": [
                    {
                        "description": "createWeeklyProgram DTO",
                        "name": "createWeeklyProgramDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateWeeklyProgramDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "returns created program's ID"
                    },
                    "400": {
                        "description": "invalid body"
                    },
                    "401": {
                        "description": "unauthorized"
                    },
                    "403": {
                        "description": "forbidden"
                    },
                    "500": {
                        "description": "failed to create program"
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateWeeklyDailyRoutineDto": {
            "type": "object",
            "properties": {
                "day": {
                    "type": "integer"
                },
                "week": {
                    "type": "integer"
                }
            }
        },
        "dto.CreateWeeklyProgramDto": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "integer"
                },
                "daily_routines": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.CreateWeeklyDailyRoutineDto"
                    }
                },
                "description": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "routine_acts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.CreateWeeklyRoutineActDto"
                    }
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "weekly_routines": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.CreateWeeklyRoutineDto"
                    }
                }
            }
        },
        "dto.CreateWeeklyRoutineActDto": {
            "type": "object",
            "properties": {
                "act_id": {
                    "type": "integer"
                },
                "day": {
                    "type": "integer"
                },
                "lap": {
                    "type": "integer"
                },
                "order": {
                    "type": "integer"
                },
                "reps": {
                    "type": "integer"
                },
                "w_ratio": {
                    "type": "number"
                },
                "warmup": {
                    "type": "boolean"
                },
                "week": {
                    "type": "integer"
                }
            }
        },
        "dto.CreateWeeklyRoutineDto": {
            "type": "object",
            "properties": {
                "week": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.0",
	Host:             "api.lifthus.com",
	BasePath:         "/routine",
	Schemes:          []string{},
	Title:            "Lifthus routine server",
	Description:      "This is Project-Hus's subservice Lifthus's routine management server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}