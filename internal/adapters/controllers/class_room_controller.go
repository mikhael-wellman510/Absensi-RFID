package controllers

import (
	"attendance-api/internal/entities"
	"attendance-api/internal/usecases"
	"attendance-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	ClassRoomController interface {
		CreateClassRoom(ctx *gin.Context)
		FindClassRoomsById(ctx *gin.Context)
	}

	classRoomController struct {
		classRoomService usecases.ClassRoomService
	}
)

func NewClassRoomController(classRoomService usecases.ClassRoomService) ClassRoomController {

	return &classRoomController{
		classRoomService: classRoomService,
	}
}

func (c *classRoomController) CreateClassRoom(ctx *gin.Context) {
	classRoomReq := entities.ClassRoomRequest{}

	if err := ctx.ShouldBind(&classRoomReq); err != nil {

		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := c.classRoomService.CreateClassRoom(ctx.Request.Context(), &classRoomReq)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.BuildResponseSuccess("Success", res))
}

func (c *classRoomController) FindClassRoomsById(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
