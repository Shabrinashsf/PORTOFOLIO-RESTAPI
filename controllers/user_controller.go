package controllers

import (
	"net/http"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/dto"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/entity"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/initializers"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/service"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	UserController interface {
		RegisterUser(ctx *gin.Context)
		Login(ctx *gin.Context)
		GetAllUser(ctx *gin.Context)
		VerifyEmail(ctx *gin.Context)
		GetUserByID(ctx *gin.Context)
	}

	userController struct {
		userService service.UserService
	}
)

func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

func (c *userController) RegisterUser(ctx *gin.Context) {
	var body dto.RegisterUserRequest

	if err := ctx.ShouldBind(&body); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response, err := c.userService.RegisterUser(ctx.Request.Context(), body)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, response)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) VerifyEmail(ctx *gin.Context) {
	code := ctx.Params.ByName("verificationCode")
	result, err := c.userService.VerifyEmail(ctx, code)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_VERIFY_EMAIL_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_VERIFY_EMAIL_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Login(ctx *gin.Context) {
	var body dto.UserLoginRequest

	if err := ctx.ShouldBind(&body); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response, err := c.userService.Verify(ctx.Request.Context(), body)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGIN, response)
	ctx.JSON(http.StatusOK, res)
}

// kayaknya gausah, tp kalo mau dibagusin sih cik
func AboutMe(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func (c *userController) GetAllUser(ctx *gin.Context) {
	users, err := c.userService.GetAllUser(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_FETCH_USERS, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_FETCH_USERS, users)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := c.userService.GetUserByID(ctx, id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_FIND_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_FIND_USER, user)
	ctx.JSON(http.StatusOK, res)
}

func UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")

	_, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid UUID format for user ID",
		})
		return
	}

	authUser, _ := ctx.Get("user")
	authUserData := authUser.(entity.User)

	if idParam != authUserData.ID.String() {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  false,
			"message": "You are not authorized to update this user",
		})
		return
	}

	var userInput struct {
		Name   string `json:"name"`
		NoTelp string `json:"no_telp"`
	}

	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid input data",
		})
		return
	}

	var user entity.User
	result := initializers.DB.First(&user, "id = ?", idParam)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "User not found",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "Failed to fetch user",
			})
		}
		return
	}

	user.Name = userInput.Name
	user.NoTelp = userInput.NoTelp

	if err := initializers.DB.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to update user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "User successfully updated",
		"user": gin.H{
			"id":      user.ID,
			"name":    user.Name,
			"no_telp": user.NoTelp,
		},
	})
}

func DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")

	_, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid UUID format for user ID",
		})
		return
	}

	var user entity.User
	result := initializers.DB.First(&user, "id = ?", idParam)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "User not found",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "Failed to fetch user",
			})
		}
		return
	}

	if err := initializers.DB.Delete(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to delete user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "User successfully deleted",
	})
}
