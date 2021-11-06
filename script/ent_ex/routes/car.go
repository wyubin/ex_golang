package routes

import (
	"example.com/ent_ex/controllers"
	"github.com/gin-gonic/gin"
)

func carRoutes(rg *gin.RouterGroup) {
	rg.POST("/car", ValidateJSONBody(controllers.CarAddReq{}), controllers.CarAdd)
}
