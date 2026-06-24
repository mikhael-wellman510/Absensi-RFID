package controllers

import (
	"attendance-api/internal/entities"
	"attendance-api/internal/usecases"
	"attendance-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	StudentParentController interface {
		CreateStudentParent(ctx *gin.Context)
		FindStudentParentById(ctx *gin.Context)
	}

	studentParentController struct {
		studentParentService usecases.StudentParentService
	}
)

func NewStudentParentController(studentParentService usecases.StudentParentService) StudentParentController {

	return &studentParentController{
		studentParentService: studentParentService,
	}
}

func (s *studentParentController) CreateStudentParent(ctx *gin.Context) {
	studentParentRequest := entities.StudentParentRequest{}

	if err := ctx.ShouldBind(&studentParentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := s.studentParentService.CreateStudentParent(ctx.Request.Context(), &studentParentRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.BuildResponseSuccess("Success", res))
}

func (s *studentParentController) FindStudentParentById(ctx *gin.Context) {

	id, ok := utils.GetIdParam(ctx, "id")

	if !ok {
		return
	}

	res, err := s.studentParentService.FindStudentParentById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Success", res))
}
