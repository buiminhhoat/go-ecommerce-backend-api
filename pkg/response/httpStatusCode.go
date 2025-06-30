package response

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 // Email is invalid
	ErrInvalidToken     = 30001 // Token is invalid
	ErrInvalidOTP       = 30002
	ErrSendEmailOtp     = 30003
	// Register Code
	ErrCodeUserHasExists = 50001 // User has alrady registered

	// Err Login
	ErrCodeOtpNotExists = 60009
)

// message
var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrInvalidOTP:        "OTP is invalid",
	ErrCodeUserHasExists: "User has alrady registered",
	ErrSendEmailOtp:      "Error when send email OTP",
	ErrCodeOtpNotExists:  "OTP exists but not registered",
}
