package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func GetPost(c *gin.Context) {
	post1 := Post{Title: "Hero", Body: "Body"}
	c.JSON(http.StatusOK, gin.H{"res": post1})
}
