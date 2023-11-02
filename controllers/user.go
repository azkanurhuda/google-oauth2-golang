package controllers

import (
	"github.com/azkanurhuda/google-oauth2-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": models.FilteredResponse(&currentUser)}})
}
