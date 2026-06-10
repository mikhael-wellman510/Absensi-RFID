package controllers

import (
	"attendance-api/internal/entities"
	"attendance-api/internal/usecases"
	"attendance-api/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	EducationLevelsController interface {
		CreateEducationLevels(ctx *gin.Context)
		FindEducationLevelsById(ctx *gin.Context)
	}

	educationLevelsController struct {
		educationLevelsService usecases.EducationLevelsService
	}
)

func NewEducationLevelsController(educationLevelsService usecases.EducationLevelsService) EducationLevelsController {

	return &educationLevelsController{
		educationLevelsService: educationLevelsService,
	}
}

func (e *educationLevelsController) CreateEducationLevels(ctx *gin.Context) {
	educationLevelsReq := entities.EducationLevelsRequest{}

	if err := ctx.ShouldBind(&educationLevelsReq); err != nil {
		log.Println("err req : ", err)
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := e.educationLevelsService.CreateEducationLevels(ctx.Request.Context(), &educationLevelsReq)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.BuildResponseSuccess("Success", res))
}

func (e *educationLevelsController) FindEducationLevelsById(ctx *gin.Context) {

	params := ctx.Param("id")

	if params == "" {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed("id is empty"))
		return
	}
	res, err := e.educationLevelsService.FindEducationLevelByID(ctx, params)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Success", res))
}
