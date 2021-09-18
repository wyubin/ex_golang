package controllers

import (
	"net/http"

	"example.com/ent_ex/models"
	"github.com/gin-gonic/gin"
)

func UserAdd(c *gin.Context) {
	user, err := models.QueryUser(c, "yubin")
	if err != nil {
		c.JSON(http.StatusInsufficientStorage, nil)
	}
	c.JSON(http.StatusOK, user)
}
