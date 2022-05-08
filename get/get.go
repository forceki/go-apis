package get

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "forceki",
		"age":  "21",
	})
}

func GetParam(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func GetQuery(c *gin.Context) {
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
