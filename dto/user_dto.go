package dto

import "errors"

const (
	// SUCCESS
	MESSAGE_SUCCESS_REGISTER_USER = "success add user"
	MESSAGE_SUCCESS_LOGIN         = "success login user"

	// FAILED
	MESSAGE_FAILED_REGISTER_USER = "failed add user"
	MESSAGE_FAILED_LOGIN         = "failed login user"
)

var (
	ErrHashPass           = errors.New("failed to hash password")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrCreateUser         = errors.New("failed to create user")
	ErrInvalidCredentials = errors.New("invalid crecentials")
	ErrAccountNotVerified = errors.New("account not verified")
	ErrFailedCreateToken  = errors.New("failed to create token")
	ErrInternalServer     = errors.New("internal server error")
)

type (
	RegisterUserRequest struct {
		Name     string `json:"name" form:"name" binding:"required"`
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
		NoTelp   string `json:"no_telp" form:"no_telp" binding:"required"`
	}

	RegisterUserResponse struct {
		Name   string `json:"name" form:"name"`
		Email  string `json:"email" form:"email"`
		NoTelp string `json:"no_telp" form:"no_telp"`
		Role   string `json:"role" form:"role"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	UserLoginResponse struct {
		Token string `json:"token" form:"token"`
		Role  string `json:"role" form:"role"`
	}
)
