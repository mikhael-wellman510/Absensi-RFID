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
	StudentController interface {
		CreateStudent(ctx *gin.Context)
		FindStudentById(ctx *gin.Context)
	}

	studentController struct {
		studentService usecases.StudentService
	}
)

func NewStudentController(studentService usecases.StudentService) StudentController {

	return &studentController{
		studentService: studentService,
	}
}

func (s *studentController) CreateStudent(ctx *gin.Context) {

	studentReq := entities.StudentRequest{}

	if err := ctx.ShouldBind(&studentReq); err != nil {
		log.Println("Invalid student request")
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := s.studentService.CreateStudent(ctx.Request.Context(), &studentReq)

	if err != nil {
		log.Println("Failed to create student")
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Success", res))
}

func (s *studentController) FindStudentById(ctx *gin.Context) {

	id, ok := utils.GetIdParam(ctx, "id")

	if !ok {
		return
	}

	res, err := s.studentService.FindStudentById(ctx.Request.Context(), id)

	if err != nil {
		log.Println("Failed to find student")
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Success", res))
}
