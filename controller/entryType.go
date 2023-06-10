package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelcclary/blogsite/model"
	"net/http"
	"strconv"
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

func GetEntryTypeByID(context *gin.Context) {
	stringId := context.Param("id")
	intId, err := strconv.Atoi(stringId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entryType, err := model.FindEntryTypeById(uint(intId))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": entryType})
}

func GetAllEntryTypes(context *gin.Context) {
	entryTypes, err := model.FindAllEntryTypes()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": entryTypes})
}
