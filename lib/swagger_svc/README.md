# intro
試著在建立 golang service 時，就可以一邊做 document，而且可以用簡單的部署，把 swagger svc 建立起來，讓 api 可以藉由 swagger 進行測試。

# code annotation
可以藉由 [swaggo](https://github.com/swaggo/swag) 來做 code 的參數註解
```golang
// @Summary Get user by ID
// @Description Get a single user by its unique ID
// @Tags user
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 404 {object} ErrorResponse
// @Router /user/{id} [get]
func GetUserByID(c *gin.Context) {
    // Your code here
}

```
以上應該只需要在 handler 的部分進行 annotation 就好

# swagger svc
可以直接用 [swagger-ui](https://github.com/swagger-api/swagger-ui/blob/HEAD/docs/usage/installation.md) 來做靜態檔服務就好．
