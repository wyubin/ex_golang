package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAdd(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
