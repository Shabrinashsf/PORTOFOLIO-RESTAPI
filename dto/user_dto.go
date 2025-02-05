package dto

import "errors"

const (
	// SUCCESS
	MESSAGE_SUCCESS_REGISTER_USER     = "success add user"
	MESSAGE_SUCCESS_LOGIN             = "success login user"
	MESSAGE_SUCCESS_FETCH_USERS       = "success to fetch users"
	MESSAGE_SUCCESS_VERIFY_EMAIL_USER = "success to verify user email verification"
	MESSAGE_SUCCESS_FIND_USER         = "success find user"

	// FAILED
	MESSAGE_FAILED_REGISTER_USER     = "failed add user"
	MESSAGE_FAILED_LOGIN             = "failed login user"
	MESSAGE_FAILED_FETCH_USERS       = "failed to fetch users"
	MESSAGE_FAILED_VERIFY_EMAIL_USER = "failed to verify user email verification"
	MESSAGE_FAILED_FIND_USER         = "failed find user"
)

var (
	ErrHashPass                = errors.New("failed to hash password")
	ErrEmailAlreadyExists      = errors.New("email already exists")
	ErrCreateUser              = errors.New("failed to create user")
	ErrInvalidCredentials      = errors.New("invalid crecentials")
	ErrAccountNotVerified      = errors.New("account not verified")
	ErrFailedCreateToken       = errors.New("failed to create token")
	ErrInternalServer          = errors.New("internal server error")
	ErrFailedGetUsers          = errors.New("failed to get users")
	ErrInvalidVerificationCode = errors.New("invalid verification code or user doesn't exists")
	ErrUserAlreadyVerified     = errors.New("user already verified")
	ErrUpdateIsVerified        = errors.New("failed to update user is_verified status")
	ErrInvalidUUID             = errors.New("invalid UUID format")
	ErrFailedFindUser          = errors.New("failed find user in database")
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

	GetAllUser struct {
		Name       string `json:"name" form:"name"`
		Email      string `json:"email" form:"email"`
		NoTelp     string `json:"no_telp" form:"no_telp"`
		Role       string `json:"role" form:"role"`
		IsVerified bool   `json:"is_verified" form:"is_verified"`
	}

	VerifyEmail struct {
		Email      string `json:"email" form:"email"`
		IsVerified bool   `json:"is_verified" form:"is_verified"`
	}

	GetUserByID struct {
		Name       string `json:"name" form:"name"`
		Email      string `json:"email" form:"email"`
		NoTelp     string `json:"no_telp" form:"no_telp"`
		Role       string `json:"role" form:"role"`
		IsVerified bool   `json:"is_verified" form:"is_verified"`
	}
)
