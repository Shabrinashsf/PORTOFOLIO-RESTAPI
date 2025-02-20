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
		routes.GET("/me", middleware.Authenticate(), controllers.AboutMe)
		routes.PUT("/user/:id", middleware.Authenticate(), userController.UpdateUser)
		routes.GET("/verifyemail/:verificationCode", userController.VerifyEmail)

		// Admin
		routes.GET("/user", middleware.Authenticate(), middleware.AdminOnly(), userController.GetAllUser)
		routes.GET("/user/:id", middleware.Authenticate(), middleware.AdminOnly(), userController.GetUserByID)
		routes.DELETE("/user/:id", middleware.Authenticate(), middleware.AdminOnly(), userController.DeleteUser)
	}
}
