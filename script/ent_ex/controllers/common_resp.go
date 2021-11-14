package controllers

// DefaultResponse 預設的回覆結果
type DefaultResponse struct {
	Status  string `example:"Success"`
	Message string
	Data    interface{} `swaggertype:"object"`
}
