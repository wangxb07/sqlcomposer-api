// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-06-01 17:33:24.528367 +0800 CST m=+0.027167309

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
        "/dbconfig": {
            "get": {
                "tags": [
                    "数据库配置"
                ],
                "summary": "数据库配置列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/restapi.DbConfigList"
                        }
                    }
                }
            },
            "patch": {
                "tags": [
                    "数据库配置"
                ],
                "summary": "添加数据库配置",
                "parameters": [
                    {
                        "description": "DbConfig",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/restapi.AddDbConfigRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "json",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/restapi.Error"
                        }
                    }
                }
            }
        },
        "/dbconfig/{uuid}": {
            "post": {
                "tags": [
                    "数据库配置"
                ],
                "summary": "更新数据库配置",
                "parameters": [
                    {
                        "description": "DbConfig",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/restapi.AddDbConfigRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "uuid",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "更新成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/restapi.Error"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "数据库配置"
                ],
                "summary": "删除数据库配置",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "json",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/restapi.Error"
                        }
                    }
                }
            }
        },
        "/doc": {
            "get": {
                "tags": [
                    "文档"
                ],
                "summary": "获取文档列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/restapi.DocListResult"
                        }
                    }
                }
            },
            "patch": {
                "tags": [
                    "文档"
                ],
                "summary": "添加新的文档",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文档内容",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "接口路径",
                        "name": "path",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "insert completed",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "deserialize yaml failed",
                        "schema": {
                            "$ref": "#/definitions/restapi.Error"
                        }
                    }
                }
            }
        },
        "/doc/{uuid}": {
            "get": {
                "tags": [
                    "文档"
                ],
                "summary": "获取文档详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Doc"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "文档"
                ],
                "summary": "更新文档",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "content",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "path",
                        "name": "path",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "update completed",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/restapi.Error"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "文档"
                ],
                "summary": "删除文档",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "delete completed",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/restapi.Error"
                        }
                    }
                }
            }
        },
        "/{path}": {
            "get": {
                "tags": [
                    "接口"
                ],
                "summary": "获取查询结果",
                "parameters": [
                    {
                        "type": "string",
                        "description": "path",
                        "name": "path",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "dbname",
                        "name": "dbname",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "json",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/restapi.Error"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/restapi.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.DataBaseConfig": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "integer"
                },
                "deleted_at": {
                    "type": "integer"
                },
                "dsn": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "integer"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "entity.Doc": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "integer"
                },
                "deleted_at": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "integer"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "restapi.AddDbConfigRequest": {
            "type": "object",
            "properties": {
                "dsn": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "restapi.DbConfigList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.DataBaseConfig"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "restapi.DocListResult": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Doc"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "restapi.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
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
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "sql-compose-api",
	Description: "This is a api for sql-compose.",
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
