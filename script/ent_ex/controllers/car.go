package controllers

import (
	"net/http"

	"example.com/ent_ex/models"
	"github.com/gin-gonic/gin"
)

// CarAdd Add car
// @Tags Car
// @Summary Add car
// @Description Add car with Model
// @Accept json
// @Produce json
// @Param request body UserQueryReq true "request"
// @Success 200 {object} string
// @Router /Car [post]
func CarAdd(c *gin.Context) {
	param, _ := c.Get("param")
	paramReq := param.(*CarAddReq)
	// check PlateNumber in db
	_, err := models.QueryCar(c, paramReq.PlateNumber)
	if err == nil {
		c.JSON(http.StatusOK, DefaultResponse{Status: "Fail", Message: "車牌已註冊"})
		return
	}
	car, err := models.CreateCar(c, paramReq.Model, paramReq.PlateNumber)
	if err != nil {
		c.JSON(http.StatusInsufficientStorage, nil)
		return
	}
	c.JSON(http.StatusOK, car)
}
