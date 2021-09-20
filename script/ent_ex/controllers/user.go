package controllers

import (
	"net/http"

	"example.com/ent_ex/models"
	"github.com/gin-gonic/gin"
)

// UserQuery 提供 User 搜尋
// @Tags User
// @Summary 提供 User 搜尋
// @Description 提供 User 搜尋
// @Accept json
// @Produce json
// @Param request body UserQueryReq true "request"
// @Success 200 {object} string
// @Router /User [get]
func UserQuery(c *gin.Context) {
	param, _ := c.Get("param")
	paramReq := param.(*UserQueryReq)
	user, err := models.QueryUser(c, paramReq.Name)
	if err != nil {
		c.JSON(http.StatusInsufficientStorage, nil)
		return
	}
	c.JSON(http.StatusOK, user)
}

// UserAdd 新增 User
// @Tags User
// @Summary 新增 User
// @Description 新增 User
// @Accept json
// @Produce json
// @Param request body UserAddReq true "request"
// @Success 200 {object} string
// @Router /User [post]
func UserAdd(c *gin.Context) {
	param, _ := c.Get("param")
	paramReq := param.(*UserAddReq)
	user, err := models.CreateUser(c, paramReq.Name, paramReq.Age)
	if err != nil {
		c.JSON(http.StatusInsufficientStorage, nil)
		return
	}
	c.JSON(http.StatusOK, user)
}
