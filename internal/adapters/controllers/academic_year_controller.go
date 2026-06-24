package controllers

import (
	"attendance-api/internal/entities"
	"attendance-api/internal/usecases"
	"attendance-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	AcademicYearController interface {
		CreateAcademicYear(ctx *gin.Context)
		FindAcademicYear(ctx *gin.Context)
	}

	academicYearController struct {
		academicYearService usecases.AcademicYearService
	}
)

func NewAcademicYearController(academicYearService usecases.AcademicYearService) AcademicYearController {

	return &academicYearController{
		academicYearService: academicYearService,
	}
}

func (a *academicYearController) CreateAcademicYear(ctx *gin.Context) {

	academicYear := entities.AcademicYearRequest{}

	if err := ctx.ShouldBind(&academicYear); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := a.academicYearService.CreateAcademicYear(ctx.Request.Context(), &academicYear)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.BuildResponseSuccess("Success", res))

}

func (a *academicYearController) FindAcademicYear(ctx *gin.Context) {

	id, ok := utils.GetIdParam(ctx, "id")

	if !ok {
		return
	}

	res, err := a.academicYearService.FindAcademicYearById(ctx.Request.Context(), id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Success", res))
}
