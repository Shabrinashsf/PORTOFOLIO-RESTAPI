package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/dto"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/initializers"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/models"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

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

		// parsing and validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_TO_PROSES_REQUEST, dto.ErrParsingToken.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// check token is valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_TO_PROSES_REQUEST, dto.ErrTokenExpired.Error(), nil)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}

			// get userID from JWT token
			userID, err := uuid.Parse(claims["user"].(string))
			if err != nil {
				response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_TO_PROSES_REQUEST, dto.ErrInvalidUserId.Error(), nil)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
			// get user role from JWT token
			userRole := claims["role"].(string)

			// look for user in db
			var user models.User
			initializers.DB.First(&user, userID)
			if user.ID == uuid.Nil {
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
