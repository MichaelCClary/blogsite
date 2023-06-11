package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelcclary/blogsite/model"
	"github.com/michaelcclary/blogsite/model/database"
	"net/http"
	"strconv"
)

func CreateEntryType(context *gin.Context) {
	var input model.EntryType
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedEntry, err := input.Create()
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

func UpdateEntryType(context *gin.Context) {
	//Verify that the entry exists
	var entryType model.EntryType
	if err := database.Database.Where("id = ?", context.Param("id")).First(&entryType).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Entry Type not found!"})
		return
	}

	//Verify the data is correct
	var input model.EntryType
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Update and return updated EntryType
	updatedEntry, err := entryType.Update(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedEntry})
}
