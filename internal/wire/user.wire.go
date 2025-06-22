//go:build wireinject

package wire

import (
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/controller"
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/repositories"
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/service"
	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repositories.NewUserRepository,
		repositories.NewUserAuthRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}
