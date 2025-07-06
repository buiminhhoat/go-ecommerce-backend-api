package account

import (
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/model"
	"github.com/buiminhhoat/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

var TwoFA = new(sUser2FA)

type sUser2FA struct{}

func (c *sUser2FA) SetupTwoFactorAuth(ctx *gin.Context) {
	var params model.SetupTwoFactorAuthInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "Missing or invalid setup two factor auth parameter")
		return
	}
	// Get UserId from uuid

}
