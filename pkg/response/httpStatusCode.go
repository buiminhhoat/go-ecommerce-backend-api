package response

const (
	ErrCodeSuccess       = 20001 // Success
	ErrCodeParamInvalid  = 20003 // Email is invalid
	ErrInvalidToken      = 30001 // Token is invalid
	ErrCodeUserHasExists = 50001 // User has alrady registered
)

// message
var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrCodeUserHasExists: "User has alrady registered",
}
