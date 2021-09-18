package routes

import (
	"example.com/ent_ex/controllers"
	"github.com/gin-gonic/gin"
)

func userRoutes(rg *gin.RouterGroup) {
	rg.GET("/user", ValidateQueryString(controllers.UserQueryReq{}), controllers.UserQuery)
	rg.POST("/user", ValidateJSONBody(controllers.UserAddReq{}), controllers.UserAdd)
}
