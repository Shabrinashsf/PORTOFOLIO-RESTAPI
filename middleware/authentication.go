package middleware

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/dto"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/entity"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/initializers"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func DecodePublicKeyBase64() ([]byte, error) {
	base64PubKey := os.Getenv("ACCESS_TOKEN_PUBLIC_KEY")
	if base64PubKey == "" {
		return nil, fmt.Errorf("public key not found in environment variables")
	}
	return base64.StdEncoding.DecodeString(base64PubKey)
}

func DecodePrivateKeyBase64() ([]byte, error) {
	base64PriKey := os.Getenv("ACCESS_TOKEN_PRIVATE_KEY")
	if base64PriKey == "" {
		return nil, fmt.Errorf("private key not found in environment variables")
	}
	return base64.StdEncoding.DecodeString(base64PriKey)
}

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get token from Authorization header
		headerAuth := ctx.GetHeader("Authorization")
		if headerAuth == "" {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_TO_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_FOUND, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// make sure token have format "Bearer <token>"
		if !strings.HasPrefix(headerAuth, "Bearer ") {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_TO_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_VALID, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// delete "Bearer " prefix to get token
		tokenString := strings.TrimPrefix(headerAuth, "Bearer ")

		// decode public key
		publicKeyBytes, err := DecodePublicKeyBase64()
		if err != nil {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_TO_PROSES_REQUEST, dto.MESSAGE_FAILED_DECODE_PUBLIC_KEY, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// parsing public key
		publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
		if err != nil {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_TO_PROSES_REQUEST, dto.MESSAGE_INVALID_PUBLIC_KEY_FORMAT, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// parsing and validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return publicKey, nil
		})
		if err != nil {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_TO_PROSES_REQUEST, dto.ErrParsingToken.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// extract claims from token
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// check expired token
			if exp, ok := claims["exp"].(float64); ok {
				if time.Now().Unix() > int64(exp) {
					response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_TO_PROSES_REQUEST, dto.ErrTokenExpired.Error(), nil)
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
					return
				}
			}

			// get user ID from claims
			userIDstr, ok := claims["user"].(string)
			if !ok {
				response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_TO_PROSES_REQUEST, dto.ErrInvalidUserId.Error(), nil)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}

			// get userID from JWT token
			userID, err := uuid.Parse(userIDstr)
			if err != nil {
				response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_TO_PROSES_REQUEST, dto.ErrInvalidUserId.Error(), nil)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}

			// get user role from JWT token
			userRole, ok := claims["role"].(string)
			if !ok {
				response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_TO_PROSES_REQUEST, dto.ErrInvalidRole.Error(), nil)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}

			// look for user in db
			var user entity.User
			result := initializers.DB.First(&user, "id = ?", userID)
			if result.Error != nil || user.ID == uuid.Nil {
				response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_TO_PROSES_REQUEST, dto.ErrUserIdEmpty.Error(), nil)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}

			// save in context
			ctx.Set("user", user)
			ctx.Set("user_id", userID.String())
			ctx.Set("role", userRole)
			ctx.Set("token", tokenString)
			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
