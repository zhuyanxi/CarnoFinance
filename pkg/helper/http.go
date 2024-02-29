package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
