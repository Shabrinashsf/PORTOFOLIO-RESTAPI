package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/constant"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/initializers"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// register, public, done
// login, public, done
// get user all, auth, by admin
// get user by id, auth, by admin

// update, auth, by user
// delete user, auth, by admin

// about me, auth

func RegisterUser(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		NoTelp   string `json:"no_telp"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Failed to hash password",
		})
		return
	}

	user := models.User{
		Name:       body.Name,
		Email:      body.Email,
		Password:   string(hash),
		NoTelp:     body.NoTelp,
		Role:       constant.ROLE_USER,
		IsVerified: false,
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "user created successfully",
		"user": gin.H{
			"name":    user.Name,
			"email":   user.Email,
			"no_telp": user.NoTelp,
		},
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Failed to read body",
			"data":    nil,
		})
		return
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == uuid.Nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Email not registered",
			"data":    nil,
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Invalid password",
			"data":    nil,
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.ID.String(),
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Failed to create token",
			"data":    nil,
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Success login",
		"data": gin.H{
			"token": tokenString,
			"role":  user.Role,
		},
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User

	if err := initializers.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Success to fetch users",
		"users":   users,
	})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid UUID format",
		})
		return
	}

	var user models.User
	result := initializers.DB.First(&user, "id = ?", parsedID)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "User not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "Failed to fetch user",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Success to fetch user",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

func UpdateUser(c *gin.Context) {
	idParam := c.Param("id")

	_, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid UUID format for user ID",
		})
		return
	}

	authUser, _ := c.Get("user")
	authUserData := authUser.(models.User)

	if idParam != authUserData.ID.String() {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  false,
			"message": "You are not authorized to update this user",
		})
		return
	}

	var userInput struct {
		Name   string `json:"name"`
		NoTelp string `json:"no_telp"`
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid input data",
		})
		return
	}

	var user models.User
	result := initializers.DB.First(&user, "id = ?", idParam)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "User not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "Failed to fetch user",
			})
		}
		return
	}

	user.Name = userInput.Name
	user.NoTelp = userInput.NoTelp

	if err := initializers.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to update user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "User successfully updated",
		"user": gin.H{
			"id":      user.ID,
			"name":    user.Name,
			"no_telp": user.NoTelp,
		},
	})
}

func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	_, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid UUID format for user ID",
		})
		return
	}

	var user models.User
	result := initializers.DB.First(&user, "id = ?", idParam)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "User not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "Failed to fetch user",
			})
		}
		return
	}

	if err := initializers.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to delete user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "User successfully deleted",
	})
}

/*// Routes
r.GET("/api/user", produkcontroller.Index)
r.GET("/api/user/:id", produkcontroller.Show)
r.POST("/api/user", produkcontroller.Create)
r.PUT("/api/user/:id", produkcontroller.Update)
r.DELETE("/api/user/:id", produkcontroller.Delete)
*/
