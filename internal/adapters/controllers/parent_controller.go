package controllers

import (
	"attendance-api/internal/entities"
	"attendance-api/internal/usecases"
	"attendance-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	ParentController interface {
		CreateParent(ctx *gin.Context)
		FindParentById(ctx *gin.Context)
	}

	parentController struct {
		parentService usecases.ParentService
	}
)

func NewParentController(parentService usecases.ParentService) ParentController {

	return &parentController{
		parentService: parentService,
	}
}

func (p *parentController) CreateParent(ctx *gin.Context) {

	parentRequest := entities.ParentRequest{}

	if err := ctx.ShouldBind(&parentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := p.parentService.CreateParent(ctx.Request.Context(), &parentRequest)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.BuildResponseSuccess("Success", res))
}

func (p *parentController) FindParentById(ctx *gin.Context) {

	id, ok := utils.GetIdParam(ctx, "id")

	if !ok {
		return
	}

	res, err := p.parentService.FindParentByID(ctx.Request.Context(), id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Success", res))

}
