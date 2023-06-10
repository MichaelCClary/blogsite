package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelcclary/blogsite/model"
	"net/http"
)

func AddEntryType(context *gin.Context) {
	var input model.EntryType
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedEntry, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}
