// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "账号登录",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_go-nunu_nunu-layout-advanced_api_v1.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_go-nunu_nunu-layout-advanced_api_v1.LoginResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "目前只支持邮箱登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_go-nunu_nunu-layout-advanced_api_v1.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_go-nunu_nunu-layout-advanced_api_v1.Response"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "获取用户信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_go-nunu_nunu-layout-advanced_api_v1.GetProfileResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "修改用户信息",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_go-nunu_nunu-layout-advanced_api_v1.UpdateProfileRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_go-nunu_nunu-layout-advanced_api_v1.GetProfileResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_go-nunu_nunu-layout-advanced_api_v1.GetProfileResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/github_com_go-nunu_nunu-layout-advanced_api_v1.GetProfileResponseData"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "github_com_go-nunu_nunu-layout-advanced_api_v1.GetProfileResponseData": {
            "type": "object",
            "properties": {
                "nickname": {
                    "type": "string",
                    "example": "alan"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "github_com_go-nunu_nunu-layout-advanced_api_v1.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "1234@gmail.com"
                },
                "password": {
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "github_com_go-nunu_nunu-layout-advanced_api_v1.LoginResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/github_com_go-nunu_nunu-layout-advanced_api_v1.LoginResponseData"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "github_com_go-nunu_nunu-layout-advanced_api_v1.LoginResponseData": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                }
            }
        },
        "github_com_go-nunu_nunu-layout-advanced_api_v1.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "1234@gmail.com"
                },
                "password": {
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "github_com_go-nunu_nunu-layout-advanced_api_v1.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "github_com_go-nunu_nunu-layout-advanced_api_v1.UpdateProfileRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "1234@gmail.com"
                },
                "nickname": {
                    "type": "string",
                    "example": "alan"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Nunu Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
