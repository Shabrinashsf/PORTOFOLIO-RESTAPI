package routes

import (
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/controllers"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/middleware"
	"github.com/gin-gonic/gin"
)

func SHSF(r *gin.Engine, shsfController controllers.SHSFController) {
	routes := r.Group("/api/shsf")
	{
		// User
		routes.POST("/register", middleware.Authenticate(), shsfController.Register)
		routes.PUT("/:id", middleware.Authenticate(), shsfController.Update)
		routes.GET("/me", middleware.Authenticate(), shsfController.GetMe)
	}
}
