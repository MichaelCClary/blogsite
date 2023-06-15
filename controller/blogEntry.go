package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelcclary/blogsite/model"
	"net/http"
	"strconv"
)

func AddBlogEntry(c *gin.Context) {
	var input model.BlogEntry
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBlogEntry, err := input.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "blog/blogEntry.html", gin.H{
		"entry": newBlogEntry,
	})
}

func ManageBlogEntry(c *gin.Context) {
	blogId := c.Param("blogId")

	var entry model.BlogEntry

	if blogId != "" {
		blogId, err := strconv.Atoi(blogId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		entry, err = model.GetBlogEntry(uint(blogId))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		entry = model.BlogEntry{}
	}

	c.HTML(http.StatusOK, "manageBlogEntry.html", gin.H{
		"entry": entry,
	})
}
