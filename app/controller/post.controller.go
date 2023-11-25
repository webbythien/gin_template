package controller

import (
	"admin-panel/v1/app/service"
	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	// Get post from mongoDB
	Post, err := service.GetPostFromMongoDB()
	if err != nil {
		c.JSON(500, gin.H{
			"status": "error",
			"data":   err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"status": "success",
			"data":   Post,
		})
	}
}
