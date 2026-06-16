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
	GradeController interface {
		CreateGrade(ctx *gin.Context)
		FindGradeById(ctx *gin.Context)
	}

	gradeController struct {
		gradeService usecases.GradeService
	}
)

func NewGradeController(gradeService usecases.GradeService) GradeController {

	return &gradeController{
		gradeService: gradeService,
	}
}

func (g *gradeController) CreateGrade(ctx *gin.Context) {
	gradeReq := entities.GradeRequest{}

	if err := ctx.ShouldBind(&gradeReq); err != nil {
		log.Println("err req : ", err)
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := g.gradeService.CreateGrade(
		ctx.Request.Context(),
		&gradeReq,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.BuildResponseSuccess("Success", res))
}

func (g *gradeController) FindGradeById(ctx *gin.Context) {

	params := ctx.Param("id")

	if params == "" {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed("id is empty"))
		return
	}

	res, err := g.gradeService.FindGradeById(
		ctx.Request.Context(),
		params,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Success", res))
}
