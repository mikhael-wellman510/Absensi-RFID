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

func (sc *schoolController) CreateSchool(ctx *gin.Context) {

	schoolReq := entities.SchoolRequest{}

	if err := ctx.ShouldBind(&schoolReq); err != nil {
		log.Println("Error request !", err.Error())
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := sc.schoolService.CreateSchool(&schoolReq)

	if err != nil {
		log.Println("Error service : ", err.Error())
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.BuildResponseSuccess("Success", res))

}
