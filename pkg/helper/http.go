package helper

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GinOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func ExtractCount(c *gin.Context) (int, error) {
	countStr := c.DefaultQuery("count", "-1")
	count, err := strconv.Atoi(countStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid count"})
		return 0, err
	}
	return count, nil
}
