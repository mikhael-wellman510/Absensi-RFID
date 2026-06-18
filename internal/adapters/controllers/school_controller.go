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
	SchoolController interface {
		CreateSchool(ctx *gin.Context)
		FindSchoolById(ctx *gin.Context)
	}

	schoolController struct {
		schoolService usecases.SchoolService
	}
)

func NewSchoolController(schoolService usecases.SchoolService) SchoolController {

	return &schoolController{
		schoolService: schoolService,
	}
}

func (s *schoolController) CreateSchool(ctx *gin.Context) {
	schoolReq := entities.SchoolRequest{}

	if err := ctx.ShouldBind(&schoolReq); err != nil {
		log.Println("err req : ", err)
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := s.schoolService.CreateSchool(
		ctx.Request.Context(),
		&schoolReq,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.BuildResponseSuccess("Success", res))
}

func (s *schoolController) FindSchoolById(ctx *gin.Context) {

	id, ok := utils.GetIdParam(ctx, "id")

	if !ok {
		return
	}

	res, err := s.schoolService.FindSchoolById(
		ctx.Request.Context(),
		id,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Success", res))
}
