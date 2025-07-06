package service

import (
	"context"

	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/model"
)

type (
	//.. interface
	IUserLogin interface {
		Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error)
		Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error)
		VerifyOTP(ctx context.Context, in *model.VerifyInput) (out model.VerifyOTPOutput, err error)
		UpdatePasswordRegister(ctx context.Context, token string, password string) (userId int, err error)

		// Two-factor authentication
		IsTwoFactorEnabled(ctx context.Context, userId int) (codeResult int, rs bool, err error)
		// Setup Authentication
		SetupTwoFactorAuth(ctx context.Context, in *model.LoginInput)
	}

	IUserInfo interface {
		GetInfoByUserId(ctx context.Context) error
		GetAllUser(ctx context.Context) error
	}

	IUserAdmin interface {
		RemoveUser(ctx context.Context) error
		FindOneUser(ctx context.Context) error
	}
)

var (
	localUserAdmin IUserAdmin
	localUserInfo  IUserInfo
	localUserLogin IUserLogin
)

func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic("Implement localUserAdmin not found for interface IUserAdmin")
	}
	return localUserAdmin
}

func InitUserAdmin(iUserAdmin IUserAdmin) {
	localUserAdmin = iUserAdmin
}

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("Implement localUserInfo not found for interface IUserInfo")
	}
	return localUserInfo
}

func InitUserLogin(iUserLogin IUserLogin) {
	localUserLogin = iUserLogin
}

func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("Implement localUserLogin not found for interface IUserLogin")
	}
	return localUserLogin
}
