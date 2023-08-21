package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func setRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		// c.String(http.StatusOK, "Hello, World!")
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Hello, World!",
		})
	})

	tes := router.Group("tes")
	tes.GET("/profile/:name", func(ctx *gin.Context) {
		param := ctx.Param("name")
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Hello, " + param,
		})
	})

	tes.POST("profile", func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid request body",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Hello, " + user.Name,
		})
	})

	tes.GET("/profile", func(ctx *gin.Context) {
		query := ctx.Query("name")
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Hello, " + query,
		})
	})

	return router
}

func main() {
	router := setRouter()
	router.Run(":8080")
}
