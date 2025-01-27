package routes

import (
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/controllers"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/middleware"
	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine, userController controllers.UserController) {
	routes := r.Group("/api")
	{
		// User
		routes.POST("/register", userController.RegisterUser)
		routes.POST("/login", userController.Login)
		routes.GET("/me", middleware.Authorization, controllers.AboutMe)
		routes.PUT("/user/:id", middleware.Authorization, controllers.UpdateUser)

		// Admin
		routes.GET("/user", middleware.Authorization, middleware.AdminOnly, controllers.GetAllUsers)
		routes.GET("/user/:id", middleware.Authorization, middleware.AdminOnly, controllers.GetUserByID)
		routes.DELETE("/user/:id", middleware.Authorization, middleware.AdminOnly, controllers.DeleteUser)
		routes.PUT("/user/validate/:id", middleware.Authorization, middleware.AdminOnly, controllers.ValidateUser)
	}
}
