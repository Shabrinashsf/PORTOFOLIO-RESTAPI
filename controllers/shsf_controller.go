package controllers

import (
	"net/http"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/dto"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/service"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/utils"
	"github.com/gin-gonic/gin"
)

type (
	SHSFController interface {
		Register(ctx *gin.Context)
		GetMe(ctx *gin.Context)
		Update(ctx *gin.Context)
	}

	shsfController struct {
		shsfService service.SHSFService
	}
)

func NewSHSFController(shsf service.SHSFService) SHSFController {
	return &shsfController{
		shsfService: shsf,
	}
}

func (c *shsfController) Register(ctx *gin.Context) {
	userIDreq, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, dto.ErrUserIdEmpty)
		return
	}

	userID := userIDreq.(string)

	var req dto.SHSFCreateRequest

	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response, err := c.shsfService.Register(ctx.Request.Context(), req, userID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_SHSF, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_SHSF, response)
	ctx.JSON(http.StatusOK, res)
}

func (c *shsfController) GetMe(ctx *gin.Context) {
	userIDreq, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, dto.ErrUserIdEmpty)
		return
	}

	response, err := c.shsfService.GetMe(ctx.Request.Context(), userIDreq)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_SHSF, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_SHSF, response)
	ctx.JSON(http.StatusOK, res)
}

func (c *shsfController) Update(ctx *gin.Context) {
	userIDreq, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, dto.ErrUserIdEmpty)
		return
	}
	userID, _ := userIDreq.(string)
	subeventID := ctx.Param("id")

	var req dto.SHSFUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response, err := c.shsfService.Update(ctx.Request.Context(), req, userID, subeventID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_SHSF, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_SHSF, response)
	ctx.JSON(http.StatusOK, res)
}
