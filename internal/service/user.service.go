package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/repositories"
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/utils/crypto"
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/utils/random"
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/utils/sendto"
	"github.com/buiminhhoat/go-ecommerce-backend-api/pkg/response"
)

// type UserService struct {
// 	userRepository *repositories.UserRepository
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepository: repositories.NewUserRepository(),
// 	}
// }

// func (us *UserService) GetInfoUser() string {
// 	return us.userRepository.GetInfoUser()
// }

// INTERFACE_VERSION

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepository     repositories.IUserRepository
	userAuthRepository repositories.IUserAuthRepository
}

func NewUserService(
	userRepository repositories.IUserRepository,
	userAuthRepository repositories.IUserAuthRepository,
) IUserService {
	return &userService{
		userRepository:     userRepository,
		userAuthRepository: userAuthRepository,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 0. hash Email
	hashEmail := crypto.GetHash(email)
	fmt.Printf("hashEmail::%s", hashEmail)
	// 5. Check OTP is available

	// 6. user spam ...

	// 1. Check email exists in db
	if us.userRepository.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	// 2. New OTP

	otp := random.GenerateSixDigitOTP()

	if purpose == "TEST_USER" {
		otp = 123456
	}

	fmt.Printf("OTP is :::%d\n", otp)

	// 3. Save OTP in Redis with expiration time
	err := us.userAuthRepository.AddOTP(email, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}
	// 4. Send Email OTP
	err = sendto.SendTemplateEmailOtp([]string{email}, "official.buiminhhoat@gmail.com",
		"otp-auth.html", map[string]interface{}{
			"otp": strconv.Itoa(otp),
		})
	if err != nil {
		return response.ErrSendEmailOtp
	}
	return response.ErrCodeSuccess
}
