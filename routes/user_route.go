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
		routes.PUT("/user/:id", middleware.Authorization, controllers.UpdateUser)

		// Admin
		routes.GET("/user", middleware.Authorization, middleware.AdminOnly, controllers.GetAllUsers)
		routes.GET("/user/:id", middleware.Authorization, middleware.AdminOnly, controllers.GetUserByID)
		routes.DELETE("/user/:id", middleware.Authorization, middleware.AdminOnly, controllers.DeleteUser)
	}
}
