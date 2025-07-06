package account

import (
	"log"

	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/model"
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/service"
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/utils/context"
	"github.com/buiminhhoat/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

var TwoFA = new(sUser2FA)

type sUser2FA struct{}

// User Step Two Factor Authentication
// @Summary      User Step Two Factor Authentication
// @Description  User Step Two Factor Authentication
// @Tags         account 2fa
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization token"
// @Param        payload body model.SetupTwoFactorAuthInput true "payload"
// @Success      200  {object} 	response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/two-factor/setup [post]
func (c *sUser2FA) SetupTwoFactorAuth(ctx *gin.Context) {
	var params model.SetupTwoFactorAuthInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "Missing or invalid setup two factor auth parameter")
		return
	}
	// Get UserId from uuid
	userId, err := context.GetUserIdFromUUID(ctx.Request.Context())
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "UserId is not valid")
	}
	log.Println("UserId: ", userId)
	params.UserId = uint32(userId)
	codeResult, err := service.UserLogin().SetupTwoFactorAuth(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeResult, nil)
}
