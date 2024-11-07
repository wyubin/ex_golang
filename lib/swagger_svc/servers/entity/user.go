package entity

type ReqUser struct {
	Name string `json:"name" example:"yubin"`
	Id   int    `json:"id" example:"1" minimum:"0"`
}
