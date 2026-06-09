package controllers

import (
	"attendance-api/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	EmailController interface {
		SendEmail(ctx *gin.Context)
	}

	emailController struct {
		emailService usecases.EmailService
	}
)

func NewEmailController(emailService usecases.EmailService) EmailController {

	return &emailController{
		emailService: emailService,
	}
}

type Bio struct {
	Name string `json:"name"`
}

func (e *emailController) SendEmail(ctx *gin.Context) {
	bio := Bio{}

	if err := ctx.ShouldBind(&bio); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := e.emailService.Email(bio.Name)

	ctx.JSON(http.StatusOK, gin.H{"result": res})
}
