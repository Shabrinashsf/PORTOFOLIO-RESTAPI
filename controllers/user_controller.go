package controllers

import (
	"net/http"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/initializers"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/models"
	"github.com/gin-gonic/gin"
)

// register, public
// login, public
// about me, auth
// update, auth
// get user all, auth
// get user by id, auth
// delete user by admin, auth

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	initializers.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
