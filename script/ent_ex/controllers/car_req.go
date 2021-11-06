package controllers

// CarAddReq CarAdd Request obj
type CarAddReq struct {
	Model string `form:"model" binding:"required"`
}
