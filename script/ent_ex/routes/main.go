package routes

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"example.com/ent_ex/configs"
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

// ValidateJSONBody 將 http req body 拆解程指定的格式
func ValidateJSONBody(reqObj interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := reflect.TypeOf(reqObj)
		body := reflect.New(t).Interface()
		RequestBody, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(RequestBody))
		err := c.ShouldBindJSON(body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(RequestBody))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		} else {
			c.Set("param", body)
			c.Next()
		}
	}
}

// ValidateQueryString 將 http query string 拆解程指定的格式
func ValidateQueryString(reqObj interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := reflect.TypeOf(reqObj)
		body := reflect.New(t).Interface()
		if err := c.ShouldBindQuery(body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		} else {
			c.Set("param", body)
			c.Next()
		}
	}
}
