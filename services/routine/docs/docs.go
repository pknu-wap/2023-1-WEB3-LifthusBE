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
        "/act": {
            "get": {
                "tags": [
                    "act"
                ],
                "summary": "gets Act name from query-string and returns corresponding Acts",
                "parameters": [
                    {
                        "type": "string",
                        "description": "act name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "skip",
                        "name": "skip",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "act id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns acts"
                    },
                    "400": {
                        "description": "invalid request"
                    },
                    "500": {
                        "description": "failed to query acts"
                    }
                }
            },
            "post": {
                "tags": [
                    "act"
                ],
                "summary": "gets CreateActDto from body and returns created act's ID",
                "parameters": [
                    {
                        "description": "createAct DTO",
                        "name": "createActDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateActDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "returns created act's ID"
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
        },
        "/program": {
            "get": {
                "tags": [
                    "program"
                ],
                "summary": "gets Program name from query-string and returns corresponding Programs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "program title",
                        "name": "title",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "skip",
                        "name": "skip",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns programs"
                    },
                    "400": {
                        "description": "invalid request"
                    },
                    "500": {
                        "description": "failed to query programs"
                    }
                }
            }
        },
        "/program/weekly": {
            "post": {
                "tags": [
                    "program"
                ],
                "summary": "gets CreateWeeklyProgramDto from body and returns created program's ID",
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
        },
        "/program/{slug}": {
            "get": {
                "tags": [
                    "program"
                ],
                "summary": "gets Program slug from path and returns corresponding Program",
                "parameters": [
                    {
                        "type": "string",
                        "description": "program slug",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns program"
                    },
                    "400": {
                        "description": "invalid request"
                    },
                    "404": {
                        "description": "program not found"
                    },
                    "500": {
                        "description": "failed to query program"
                    }
                }
            }
        },
        "/rec/program/weekly": {
            "post": {
                "tags": [
                    "rec"
                ],
                "summary": "gets CreateWeeklyProgramRecDto from body and returns created rec's ID",
                "parameters": [
                    {
                        "description": "createWeeklyProgramRec DTO",
                        "name": "createWeeklyProgramRecDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateWeeklyProgramRecDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "returns created rec's ID"
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
                        "description": "failed to create rec"
                    }
                }
            }
        },
        "/rec/routineact": {
            "get": {
                "tags": [
                    "rec"
                ],
                "summary": "gets specific date or range of date and returns routineact recs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "date like 2006-01-02",
                        "name": "date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "start date like 2006-01-02",
                        "name": "startDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "end date like 2006-01-02s",
                        "name": "endDate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns routineact recs"
                    },
                    "400": {
                        "description": "invalid request"
                    },
                    "401": {
                        "description": "unauthorized"
                    },
                    "403": {
                        "description": "forbidden"
                    },
                    "500": {
                        "description": "failed to query routineact recs"
                    }
                }
            },
            "put": {
                "tags": [
                    "rec"
                ],
                "summary": "updates routineact rec",
                "parameters": [
                    {
                        "description": "update routineact rec dto",
                        "name": "updateRoutineActRecDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateRoutineActRecDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns updated routineact rec"
                    },
                    "400": {
                        "description": "invalid request"
                    },
                    "401": {
                        "description": "unauthorized"
                    },
                    "403": {
                        "description": "forbidden"
                    },
                    "500": {
                        "description": "failed to update routineact rec"
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateActDto": {
            "type": "object",
            "properties": {
                "arms": {
                    "description": "prime movers",
                    "type": "boolean"
                },
                "author": {
                    "type": "integer"
                },
                "bodyweight": {
                    "type": "boolean"
                },
                "cardio": {
                    "type": "boolean"
                },
                "chest": {
                    "type": "boolean"
                },
                "core": {
                    "type": "boolean"
                },
                "description": {
                    "type": "string"
                },
                "etc": {
                    "type": "boolean"
                },
                "full": {
                    "type": "boolean"
                },
                "glute": {
                    "type": "boolean"
                },
                "image": {
                    "type": "string"
                },
                "legs_back": {
                    "type": "boolean"
                },
                "legs_front": {
                    "type": "boolean"
                },
                "lower": {
                    "type": "boolean"
                },
                "lower_back": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "shoulders": {
                    "type": "boolean"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "type": {
                    "type": "string"
                },
                "upper": {
                    "description": "upper/lower/full body",
                    "type": "boolean"
                },
                "upper_back": {
                    "type": "boolean"
                },
                "weight": {
                    "description": "weight/cardio",
                    "type": "boolean"
                }
            }
        },
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
        "dto.CreateWeeklyDailyRoutineRecDto": {
            "type": "object",
            "properties": {
                "daily_routine_id": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "routine_act_recs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.CreateWeeklyRoutineActRecDto"
                    }
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
        "dto.CreateWeeklyProgramRecDto": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "integer"
                },
                "comment": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "program_id": {
                    "type": "integer"
                },
                "start_date": {
                    "type": "string"
                },
                "weekly_routine_recs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.CreateWeeklyRoutineRecDto"
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
        "dto.CreateWeeklyRoutineActRecDto": {
            "type": "object",
            "properties": {
                "act_id": {
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
                "routine_act_id": {
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
        },
        "dto.CreateWeeklyRoutineRecDto": {
            "type": "object",
            "properties": {
                "daily_routine_recs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.CreateWeeklyDailyRoutineRecDto"
                    }
                },
                "start_date": {
                    "type": "string"
                },
                "week": {
                    "type": "integer"
                },
                "weekly_routine_id": {
                    "type": "integer"
                }
            }
        },
        "dto.UpdateRoutineActRecDto": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "integer"
                },
                "comment": {
                    "type": "string"
                },
                "current_lap": {
                    "type": "integer"
                },
                "current_reps": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "started_at": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
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
