package main

import (
	"log"
	"net/http"

	"go-apis/get"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", get.GetTest)
	router.GET("/param/:id", get.GetParam)
	router.GET("/query", get.GetQuery)
	router.POST("/post", postTest)

	router.Run()
}

type TestPost struct {
	Text  string
	Title string
}

func postTest(c *gin.Context) {
	var testPost TestPost

	err := c.ShouldBindJSON(&testPost)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"title": testPost.Title,
		"text":  testPost.Text,
	})
}
