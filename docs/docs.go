// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-05-23 17:24:34.8450535 +0700 +07 m=+0.069982701

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
        "termsOfService": "COJ",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "เข้าสู่ระบบ",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "เข้าสู่ระบบ"
                ],
                "summary": "เข้าสู่ระบบ",
                "parameters": [
                    {
                        "description": "ลงชื่อเข้าสู่ระบบ",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SwagLogin"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "รายการผู้ใช้งาน",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ผู้ใช้งาน"
                ],
                "summary": "รายการผู้ใช้งาน",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SwagGetAllUsersResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.SwagError400"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.SwagError404"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.SwagError500"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "เพิ่มผู้ใช้งาน",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ผู้ใช้งาน"
                ],
                "summary": "เพิ่มผู้ใช้งาน",
                "parameters": [
                    {
                        "description": "เพิ่มผู้ใช้งาน",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SwagUserBodyIncludePassword"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.SwagCreateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.SwagError400"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.SwagError404"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.SwagError500"
                        }
                    }
                }
            }
        },
        "/users/{slug}": {
            "get": {
                "description": "ข้อมูลผู้ใช้งาน",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ผู้ใช้งาน"
                ],
                "summary": "ข้อมูลผู้ใช้งาน",
                "parameters": [
                    {
                        "type": "string",
                        "description": "slug ผู้ใช้งาน",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SwagGetUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.SwagError400"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.SwagError404"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.SwagError500"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "แก้ไขผู้ใช้งาน",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ผู้ใช้งาน"
                ],
                "summary": "แก้ไขผู้ใช้งาน",
                "parameters": [
                    {
                        "type": "string",
                        "description": "slug ผู้ใช้งาน",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "แก้ไขผู้ใช้งาน",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SwagUserBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.SwagUpdateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.SwagError400"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.SwagError404"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.SwagError500"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "ลบผู้ใช้งาน",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ผู้ใช้งาน"
                ],
                "summary": "ลบผู้ใช้งาน",
                "parameters": [
                    {
                        "type": "string",
                        "description": "slug ผู้ใช้งาน",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.SwagDeleteBase"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.SwagError400"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.SwagError404"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.SwagError500"
                        }
                    }
                }
            }
        },
        "/users/{slug}/password": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "เปลี่ยนรหัสผ่าน",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ผู้ใช้งาน"
                ],
                "summary": "เปลี่ยนรหัสผ่าน",
                "parameters": [
                    {
                        "type": "string",
                        "description": "slug ผู้ใช้งาน",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "เปลี่ยนรหัสผ่าน",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ChangePassword"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.SwagChangePasswordResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.SwagError400"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.SwagError404"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.SwagError500"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ChangePassword": {
            "type": "object",
            "required": [
                "current_password",
                "new_password"
            ],
            "properties": {
                "current_password": {
                    "description": "รหัสผ่านปัจจุบัน",
                    "type": "string",
                    "example": "password"
                },
                "new_password": {
                    "description": "รหัสผ่านใหม่",
                    "type": "string",
                    "example": "password123"
                }
            }
        },
        "models.SwagChangePasswordResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "ข้อมูล",
                    "type": "object"
                },
                "message": {
                    "description": "ข้อความตอบกลับ",
                    "type": "string",
                    "example": "Change Password Successfully"
                },
                "status": {
                    "description": "HTTP Status Code",
                    "type": "integer",
                    "example": 200
                },
                "success": {
                    "description": "ผลการเรียกใช้งาน",
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.SwagCreateUserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "users": {
                            "type": "object",
                            "$ref": "#/definitions/models.SwagUser"
                        }
                    }
                },
                "message": {
                    "description": "ข้อความตอบกลับ",
                    "type": "string",
                    "example": "Created Successfully"
                },
                "status": {
                    "description": "HTTP Status Code",
                    "type": "integer",
                    "example": 201
                },
                "success": {
                    "description": "ผลการเรียกใช้งาน",
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.SwagDeleteBase": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "ข้อมูล",
                    "type": "object"
                },
                "message": {
                    "description": "ข้อความตอบกลับ",
                    "type": "string",
                    "example": "Deleted Successfully"
                },
                "status": {
                    "description": "HTTP Status Code",
                    "type": "integer",
                    "example": 200
                },
                "success": {
                    "description": "ผลการเรียกใช้งาน",
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.SwagError400": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "ข้อมูล",
                    "type": "object"
                },
                "message": {
                    "description": "ข้อความตอบกลับ",
                    "type": "string",
                    "example": "Bad Request"
                },
                "status": {
                    "description": "HTTP Status Code",
                    "type": "integer",
                    "example": 400
                },
                "success": {
                    "description": "ผลการเรียกใช้งาน",
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "models.SwagError404": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "ข้อมูล",
                    "type": "object"
                },
                "message": {
                    "description": "ข้อความตอบกลับ",
                    "type": "string",
                    "example": "Not Found"
                },
                "status": {
                    "description": "HTTP Status Code",
                    "type": "integer",
                    "example": 404
                },
                "success": {
                    "description": "ผลการเรียกใช้งาน",
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "models.SwagError500": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "ข้อมูล",
                    "type": "object"
                },
                "message": {
                    "description": "ข้อความตอบกลับ",
                    "type": "string",
                    "example": "Internal Server Error"
                },
                "status": {
                    "description": "HTTP Status Code",
                    "type": "integer",
                    "example": 500
                },
                "success": {
                    "description": "ผลการเรียกใช้งาน",
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "models.SwagGetAllUsersResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "pageMeta": {
                            "type": "object",
                            "$ref": "#/definitions/models.SwagPageMeta"
                        },
                        "users": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.SwagUser"
                            }
                        }
                    }
                },
                "message": {
                    "description": "ข้อความตอบกลับ",
                    "type": "string",
                    "example": "Data Found"
                },
                "status": {
                    "description": "HTTP Status Code",
                    "type": "integer",
                    "example": 200
                },
                "success": {
                    "description": "ผลการเรียกใช้งาน",
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.SwagGetUserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "users": {
                            "type": "object",
                            "$ref": "#/definitions/models.SwagUser"
                        }
                    }
                },
                "message": {
                    "description": "ข้อความตอบกลับ",
                    "type": "string",
                    "example": "Data Found"
                },
                "status": {
                    "description": "HTTP Status Code",
                    "type": "integer",
                    "example": 200
                },
                "success": {
                    "description": "ผลการเรียกใช้งาน",
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.SwagLogin": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "$ref": "#/definitions/models.SwagUser"
                },
                "message": {
                    "description": "ข้อความตอบกลับ",
                    "type": "string",
                    "example": "User logged in successfully"
                },
                "status": {
                    "description": "HTTP Status Code",
                    "type": "integer",
                    "example": 200
                },
                "success": {
                    "description": "ผลการเรียกใช้งาน",
                    "type": "boolean",
                    "example": true
                },
                "token": {
                    "description": "token",
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsI"
                }
            }
        },
        "models.SwagPageMeta": {
            "type": "object",
            "properties": {
                "currentItemsCount": {
                    "type": "integer",
                    "example": 1
                },
                "currentPageNumber": {
                    "type": "integer",
                    "example": 1
                },
                "hasNextPage": {
                    "type": "boolean",
                    "example": false
                },
                "hasPrevPage": {
                    "type": "boolean",
                    "example": false
                },
                "nextPageNumber": {
                    "type": "integer",
                    "example": 1
                },
                "nextPageUrl": {
                    "type": "string",
                    "example": "/api/v1/users/addresses?page=1\u0026pageSize=10"
                },
                "offset": {
                    "type": "integer",
                    "example": 0
                },
                "prevPageNumber": {
                    "type": "integer",
                    "example": 1
                },
                "prevPageUrl": {
                    "type": "string",
                    "example": "/api/v1/users/addresses?page=1\u0026pageSize=10"
                },
                "requestedPageSize": {
                    "type": "integer",
                    "example": 10
                },
                "totalItemsCount": {
                    "type": "integer",
                    "example": 1
                },
                "totalPagesCount": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "models.SwagUpdateUserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "users": {
                            "type": "object",
                            "$ref": "#/definitions/models.SwagUser"
                        }
                    }
                },
                "message": {
                    "description": "ข้อความตอบกลับ",
                    "type": "string",
                    "example": "Updated Successfully"
                },
                "status": {
                    "description": "HTTP Status Code",
                    "type": "integer",
                    "example": 200
                },
                "success": {
                    "description": "ผลการเรียกใช้งาน",
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.SwagUser": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "รูป Avatar",
                    "type": "string",
                    "example": "user01.png"
                },
                "email": {
                    "description": "อีเมล์",
                    "type": "string",
                    "example": "user01@email.com"
                },
                "firstName": {
                    "description": "ชื่อ",
                    "type": "string",
                    "example": "john"
                },
                "id": {
                    "description": "ID",
                    "type": "integer",
                    "example": 1
                },
                "lastName": {
                    "description": "นามสกุล",
                    "type": "string",
                    "example": "doe"
                },
                "slug": {
                    "description": "Slug",
                    "type": "string",
                    "example": "user01"
                },
                "username": {
                    "description": "Username",
                    "type": "string",
                    "example": "user01"
                }
            }
        },
        "models.SwagUserBody": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "รูป Avatar",
                    "type": "string",
                    "example": "user01.png"
                },
                "email": {
                    "description": "อีเมล์",
                    "type": "string",
                    "example": "user01@email.com"
                },
                "firstName": {
                    "description": "ชื่อ",
                    "type": "string",
                    "example": "john"
                },
                "lastName": {
                    "description": "นามสกุล",
                    "type": "string",
                    "example": "doe"
                },
                "slug": {
                    "description": "Slug",
                    "type": "string",
                    "example": "user01"
                },
                "username": {
                    "description": "Username",
                    "type": "string",
                    "example": "user01"
                }
            }
        },
        "models.SwagUserBodyIncludePassword": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "รูป Avatar",
                    "type": "string",
                    "example": "user01.png"
                },
                "email": {
                    "description": "อีเมล์",
                    "type": "string",
                    "example": "user01@email.com"
                },
                "firstName": {
                    "description": "ชื่อ",
                    "type": "string",
                    "example": "john"
                },
                "lastName": {
                    "description": "นามสกุล",
                    "type": "string",
                    "example": "doe"
                },
                "password": {
                    "description": "รหัสผ่าน",
                    "type": "string",
                    "example": "pass1234"
                },
                "slug": {
                    "description": "Slug",
                    "type": "string",
                    "example": "user01"
                },
                "username": {
                    "description": "Username",
                    "type": "string",
                    "example": "user01"
                }
            }
        },
        "routes.login": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Example API",
	Description: "Rest API document <style>.models {display: none !important}</style>",
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