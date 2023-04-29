// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/order/items": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Create Order",
                "operationId": "createorder",
                "parameters": [
                    {
                        "description": "Order Detials",
                        "name": "orderdetials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_SA01-OrderManagement_SA01-grpc-api-gateway_pkg_domain.Order"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_SA01-OrderManagement_SA01-grpc-api-gateway_pkg_utils_response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_SA01-OrderManagement_SA01-grpc-api-gateway_pkg_utils_response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_fazilnbr_SA01-OrderManagement_SA01-grpc-api-gateway_pkg_domain.Item": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "github_com_fazilnbr_SA01-OrderManagement_SA01-grpc-api-gateway_pkg_domain.Order": {
            "type": "object",
            "properties": {
                "currencyUnit": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_fazilnbr_SA01-OrderManagement_SA01-grpc-api-gateway_pkg_domain.Item"
                    }
                },
                "status": {
                    "type": "string"
                },
                "total": {
                    "type": "number"
                }
            }
        },
        "github_com_fazilnbr_SA01-OrderManagement_SA01-grpc-api-gateway_pkg_utils_response.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "errors": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}