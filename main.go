package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyper-dot/bank-management-system/router"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Home Page",
		})
	})
	r.GET("/post", router.GetPost)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":3000") // Listens on port 3000
}
