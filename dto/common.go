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
	MESSAGE_FAILED_DECODE_PUBLIC_KEY  = "failed to decode public key"
	MESSAGE_INVALID_PUBLIC_KEY_FORMAT = "Invalid public key format"
)

var (
	ErrUserIdEmpty             = errors.New("empty user ID")
	ErrTokenExpired            = errors.New("token expired")
	ErrParsingToken            = errors.New("failed parsing token")
	ErrInvalidUserId           = errors.New("invalid user id in token")
	ErrInvalidRole             = errors.New("invalid role")
	ErrFailedDecodePrivateKey  = errors.New("failed to decode private key")
	ErrInvalidPrivateKeyFormat = errors.New("invalid private key format")
	ErrUserNotFound            = errors.New("user not found")
	ErrGeneral                 = errors.New("something went wrong")
	ErrGetUserById             = errors.New("failed get user by id")
)
