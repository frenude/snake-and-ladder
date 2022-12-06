package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, Welcome to GIN FrameWork")
}
