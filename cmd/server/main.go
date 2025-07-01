package main

import (
	_ "github.com/buiminhhoat/go-ecommerce-backend-api/cmd/swag/docs"
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/initialize"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API Documentation Ecommerce Backend SHOPDEVGO
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  https://github.com/buiminhhoat/go-ecommerce-backend-api

// @contact.name   Hoat Bui Minh
// @contact.url    https://github.com/buiminhhoat/go-ecommerce-backend-api
// @contact.email  official.buiminhhoat@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8888
// @BasePath  /v1/2025
// @schemes   http

func main() {
	r := initialize.Run()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8888")
}
