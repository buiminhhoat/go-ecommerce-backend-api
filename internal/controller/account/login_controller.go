package account

import (
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/service"
	"github.com/buiminhhoat/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type cUserLogin struct{}

var Login = new(cUserLogin)

func (c *cUserLogin) Login(ctx *gin.Context) {
	// Implement logic for login

	err := service.UserLogin().Login(ctx)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, nil)
}
