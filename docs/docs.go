// Code generated by swaggo/swag. DO NOT EDIT
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
        "/session/new": {
            "post": {
                "description": "at the same time user connects to lifthus newly, the client requests new session token.\nand the server returns session id with session token in cookie.",
                "tags": [
                    "auth"
                ],
                "summary": "gets new connection and assign a lifthus session token.",
                "responses": {
                    "201": {
                        "description": "returns session id with session token in cookie"
                    },
                    "501": {
                        "description": "failed to create new session"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.0",
	Host:             "lifthus.com",
	BasePath:         "/auth",
	Schemes:          []string{},
	Title:            "Project-Hus auth server",
	Description:      "This is Project-Hus's root authentication server containing each user's UUID, which is unique for all hus services.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}