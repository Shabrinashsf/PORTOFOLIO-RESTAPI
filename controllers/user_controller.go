package controllers

import (
	"net/http"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/dto"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/entity"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/service"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/utils"
	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		RegisterUser(ctx *gin.Context)
		Login(ctx *gin.Context)
		GetAllUser(ctx *gin.Context)
		VerifyEmail(ctx *gin.Context)
		GetUserByID(ctx *gin.Context)
		UpdateUser(ctx *gin.Context)
		DeleteUser(ctx *gin.Context)
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

func (c *userController) UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, "ID is required", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var body dto.UpdateUser
	if err := ctx.ShouldBindJSON(&body); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response, err := c.userService.UpdateUser(ctx, idParam, body)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_USER, response)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, "ID is required", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var user entity.User
	response, err := c.userService.DeleteUser(idParam, user)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_USER, response)
	ctx.JSON(http.StatusOK, res)
}
