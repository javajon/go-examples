package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func startGin() {
	router := gin.Default()

	// Register handlers
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Home page")
	})

	// Idea CRUD routes and handlers
	ideaGroup := router.Group("ideas/v1")
	{
		ideaGroup.GET("/", GetAll)
		ideaGroup.GET("/:id", Get)
		ideaGroup.POST("/", Create)
		ideaGroup.PATCH("/:id", Update)
		ideaGroup.DELETE("/:id", Delete)
	}

	router.Run(":8080")
}

func main() {
	startGin()
}
