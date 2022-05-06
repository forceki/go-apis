package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", getTest)
	router.GET("/param/:id", getParam)
	router.GET("/query", getQuery)

	router.Run()
}

func getTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "forceki",
		"age":  "21",
	})
}

func getParam(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func getQuery(c *gin.Context) {
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
