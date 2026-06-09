package controllers

import (
	"attendance-api/internal/entities"
	"attendance-api/internal/usecases"
	"attendance-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	TeachersController interface {
		CreateTeachers(ctx *gin.Context)
	}

	teachersController struct {
		teachersService usecases.TeachersService
	}
)

func NewTeachersController(teachersService usecases.TeachersService) TeachersController {

	return &teachersController{
		teachersService: teachersService,
	}
}

func (tc *teachersController) CreateTeachers(ctx *gin.Context) {

	teachersReq := entities.TeachersRequest{}

	if err := ctx.ShouldBind(&teachersReq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := tc.teachersService.CreateTeachers(ctx.Request.Context(), &teachersReq)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Success", res))

}
