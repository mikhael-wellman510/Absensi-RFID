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
	UserController interface {
		CreateUser(ctx *gin.Context)
		FindUserById(ctx *gin.Context)
	}

	userController struct {
		userService usecases.UserService
	}
)

func NewUserController(userService usecases.UserService) UserController {

	return &userController{
		userService: userService,
	}
}

func (u *userController) CreateUser(ctx *gin.Context) {
	userReq := entities.UserRequest{}

	if err := ctx.ShouldBind(&userReq); err != nil {
		log.Println("err req : ", err)
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := u.userService.CreateUser(
		ctx.Request.Context(),
		&userReq,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.BuildResponseSuccess("Success", res))
}

func (u *userController) FindUserById(ctx *gin.Context) {

	params := ctx.Param("id")

	if params == "" {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed("id is empty"))
		return
	}

	res, err := u.userService.FindUserById(
		ctx.Request.Context(),
		params,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Success", res))
}
