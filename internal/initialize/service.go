package initialize

import (
	"github.com/buiminhhoat/go-ecommerce-backend-api/global"
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/database"
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/service"
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/service/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)
	// User Service Interface
	service.InitUserLogin(impl.NewUserLoginImpl(queries))
	// ...
}
