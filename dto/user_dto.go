package dto

import "errors"

const (
	// SUCCESS
	MESSAGE_SUCCESS_REGISTER_USER = "success add user"

	// FAILED
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_FAILED_REGISTER_USER      = "failed add user"
)

var (
	ErrHashPass           = errors.New("failed to hash password")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrCreateUser         = errors.New("failed to create user")
)

type (
	RegisterUserRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		NoTelp   string `json:"no_telp"`
	}

	RegisterUserResponse struct {
		Name   string `json:"name"`
		Email  string `json:"email"`
		NoTelp string `json:"no_telp"`
		Role   string `json:"role"`
	}
)
