package dto

import "errors"

const (
	// FAILED
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_FAILED_TO_PROSES_REQUEST  = "failed to proses request"
	MESSAGE_FAILED_TOKEN_NOT_FOUND    = "failed to found token"
	MESSAGE_FAILED_TOKEN_NOT_VALID    = "failed invalid token"
	MESSAGE_FAILED_ACCESS_DENIED      = "failed access for non admin"
	MESSAGE_FAILED_USER_NOT_FOUND     = "failed access user not found"
)

var (
	ErrUserIdEmpty   = errors.New("empty user ID")
	ErrTokenExpired  = errors.New("token expired")
	ErrParsingToken  = errors.New("failed parsing token")
	ErrInvalidUserId = errors.New("invalid user id in token")
)
