# intro
試著在建立 golang service 時，就可以一邊做 document，而且可以用簡單的部署，把 swagger svc 建立起來，讓 api 可以藉由 swagger 進行測試。

# requirement
## install
```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

## code annotation
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

## swagger svc
可以直接用 [swagger-ui](https://github.com/swagger-api/swagger-ui/blob/HEAD/docs/usage/installation.md) 來做靜態檔服務就好．

# plan
可以先試著建一個簡單的 svc 來試試，首先就要可以 service static 跟一些 api
基本route 架構 - 可以參考[route.go](https://github.com/benhoyt/go-routing/blob/master/stdlib/route.go)

## test
```shell
# get route info
curl localhost:3000/info/
# post user
curl -X POST localhost:3000/info/user \
    -d '{"name":"wyubin", "id": 1}'
```

# docs
## install ui
把 import unpkg 的 [html](https://github.com/swagger-api/swagger-ui/blob/HEAD/docs/usage/installation.md#unpkg) 放到 static 內就可以，加上設定一個內部產生的 swagger.yaml

## generate
```shell
dir_static=./static
swag init \
    --output ${dir_static} \
    --outputTypes yaml
```

## note
要注意的可能是進行 k8s 部屬時，可能要注意要讓 docs 打的 host 是可以依照 env 改變的，所以可能在環境變數上要給一個 SWAGGER_API_HOST 讓 golang 裡面可以 dynamic 寫入這個變數

也可以直接在 html 裡面做動態修正
```javascript
window.onload = () => {
  const basePath = window.location.pathname.substring(0, window.location.pathname.indexOf('/static/') + 1);
  window.ui = SwaggerUIBundle({
    url: './swagger.yaml',
    dom_id: '#swagger-ui',
    host: window.location.host,
    basePath: basePath,
  });
};
```