package middleware

import (
	"net/http"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/dto"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/utils"
	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exists := ctx.Get("role")
		if !exists || role.(string) != "admin" {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_TO_PROSES_REQUEST, dto.MESSAGE_FAILED_ACCESS_DENIED, nil)
			ctx.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}

		user, exists := ctx.Get("user")
		if !exists {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_TO_PROSES_REQUEST, dto.MESSAGE_FAILED_USER_NOT_FOUND, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
