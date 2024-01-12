package utilites

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AbortWithErrorMessage(ctx *gin.Context, message string) {
	errorMessage := new(models.Error)
	errorMessage.Message = message
	ctx.JSON(http.StatusBadRequest, errorMessage)
}
