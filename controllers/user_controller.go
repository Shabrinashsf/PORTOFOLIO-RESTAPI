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
// login, public
// about me, auth
// update, auth, done
// get user all, auth, done
// get user by id, auth, done
// delete user by admin, auth

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

func GetAllUser(c *gin.Context) {

	var user []models.User

	initializers.DB.Find(&user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetUserByID(c *gin.Context) {

	var user models.User
	id := c.Param("id")

	if err := initializers.DB.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {

	var produk models.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&produk); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if initializers.DB.Model(&produk).Where("id = ?", id).Updates(&produk).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbaharui"})
}

func DeleteUser(c *gin.Context) {

	var user models.User

	id := c.Param("id")

	if err := initializers.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Record not found!"})
		return
	}

	if initializers.DB.Delete(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}

/*// Routes
r.GET("/api/user", produkcontroller.Index)
r.GET("/api/user/:id", produkcontroller.Show)
r.POST("/api/user", produkcontroller.Create)
r.PUT("/api/user/:id", produkcontroller.Update)
r.DELETE("/api/user/:id", produkcontroller.Delete)
*/
