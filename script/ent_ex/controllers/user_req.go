package controllers

// UserQueryReq UserQuery Request obj
type UserQueryReq struct {
	Name string `form:"name" binding:"required"`
}

// UserQueryReq UserQuery Request obj
type UserAddReq struct {
	Name string `binding:"required"`
	Age  *int
}
