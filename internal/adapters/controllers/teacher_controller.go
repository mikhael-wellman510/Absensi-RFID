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
	TeacherController interface {
		CreateTeacher(ctx *gin.Context)
		FindTeacherById(ctx *gin.Context)
	}

	teacherController struct {
		teacherService usecases.TeacherService
	}
)

func NewTeacherController(teacherService usecases.TeacherService) TeacherController {

	return &teacherController{
		teacherService: teacherService,
	}
}

func (t *teacherController) CreateTeacher(ctx *gin.Context) {

	teacherReq := entities.TeacherRequest{}

	if err := ctx.ShouldBind(&teacherReq); err != nil {
		log.Println("invalid teacher request")
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := t.teacherService.CreateTeacher(ctx.Request.Context(), &teacherReq)

	if err != nil {

		log.Println("Failed to create teacher")
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Success", res))
}

func (t *teacherController) FindTeacherById(ctx *gin.Context) {

	id, ok := utils.GetIdParam(ctx, "id")

	if !ok {
		return
	}

	res, err := t.teacherService.FindTeacherById(ctx.Request.Context(), id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Success", res))
}
