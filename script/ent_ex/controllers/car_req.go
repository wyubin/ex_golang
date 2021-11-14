package controllers

// CarAddReq CarAdd Request obj
type CarAddReq struct {
	Model       string `binding:"required"`
	PlateNumber string `binding:"required"`
}
