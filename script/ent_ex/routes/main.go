package routes

import (
	"fmt"

	"example.com/ent_ex/configs"
	"example.com/ent_ex/controllers"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.New()
)

func Run() {
	rg := router.Group("/")
	userRoutes(rg)
	_ = router.Run(fmt.Sprintf(":%d", configs.Cfg.Gin.Port))
}

func userRoutes(rg *gin.RouterGroup) {
	rg.GET("/user", controllers.UserAdd)
}
