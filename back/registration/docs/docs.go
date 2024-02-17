// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/register": {
            "post": {
                "description": "Добавляет новую компанию и её владельца, который является её первым сотрудником, одним запросом",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Добавление новой компании и владельца",
                "parameters": [
                    {
                        "description": "Информация о компании и её владельце",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpserver.addCompanyAndOwnerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное добавление компании с владельцем",
                        "schema": {
                            "$ref": "#/definitions/httpserver.companyAndOwnerResponse"
                        }
                    },
                    "400": {
                        "description": "Неверный формат входных данных",
                        "schema": {
                            "$ref": "#/definitions/httpserver.companyAndOwnerResponse"
                        }
                    },
                    "500": {
                        "description": "Проблемы на стороне сервера",
                        "schema": {
                            "$ref": "#/definitions/httpserver.companyAndOwnerResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httpserver.addCompanyAndOwnerRequest": {
            "type": "object",
            "properties": {
                "company": {
                    "$ref": "#/definitions/httpserver.addCompanyData"
                },
                "owner": {
                    "$ref": "#/definitions/httpserver.addOwnerData"
                }
            }
        },
        "httpserver.addCompanyData": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "industry": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "httpserver.addOwnerData": {
            "type": "object",
            "properties": {
                "department": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "job_title": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "second_name": {
                    "type": "string"
                }
            }
        },
        "httpserver.companyAndOwnerData": {
            "type": "object",
            "properties": {
                "company": {
                    "$ref": "#/definitions/httpserver.companyData"
                },
                "owner": {
                    "$ref": "#/definitions/httpserver.ownerData"
                }
            }
        },
        "httpserver.companyAndOwnerResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/httpserver.companyAndOwnerData"
                },
                "err": {
                    "type": "string"
                }
            }
        },
        "httpserver.companyData": {
            "type": "object",
            "properties": {
                "creation_date": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "industry": {
                    "type": "integer"
                },
                "is_deleted": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                }
            }
        },
        "httpserver.ownerData": {
            "type": "object",
            "properties": {
                "company_id": {
                    "type": "integer"
                },
                "creation_date": {
                    "type": "integer"
                },
                "department": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_deleted": {
                    "type": "boolean"
                },
                "job_title": {
                    "type": "string"
                },
                "second_name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8091",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "BRM API",
	Description:      "Swagger документация к API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
