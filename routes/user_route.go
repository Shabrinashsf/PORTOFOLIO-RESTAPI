package routes

import (
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/controllers"
	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine) {
	routes := route.Group("/api/auth")
	{
		// User
		routes.POST("", controllers.Register)
		//routes.POST("/login", userController.Login)
		//routes.PUT("/update", middleware.Authenticate(jwtService), userController.Update)
		//routes.POST("/sendmail", userController.SendVerificationEmail)
		//routes.GET("", middleware.Authenticate(jwtService), middleware.OnlyAllow("admin"), userController.GetAllUser)
		//routes.GET("/verify", userController.VerifyEmail)
		//routes.GET("/me", middleware.Authenticate(jwtService), userController.Me)
		//routes.POST("/reset", userController.ResetPassword)
		//routes.POST("/forget", userController.ForgetPassword)

	}
}
