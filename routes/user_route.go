package routes

import (
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/controllers"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/middleware"
	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine) {
	routes := r.Group("/api")
	{
		// User
		routes.POST("/register", controllers.RegisterUser)
		routes.POST("/login", controllers.Login)
		routes.GET("/validate", middleware.Authorization, controllers.Validate)
	}
}
