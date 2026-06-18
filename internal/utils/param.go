package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIdParam(ctx *gin.Context, paramName string) (string, bool) {

	param := ctx.Param(paramName)

	if param == "" {
		ctx.JSON(
			http.StatusBadRequest, BuildResponseFailed(param+" is empty"))
		return "", false
	}

	return param, true
}
