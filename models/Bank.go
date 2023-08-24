package models

type Customer struct {
	Name     string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
}
type UpdateName struct{
	IntialName string `json:"intialname" bson:"intialname"`
	UpdateName string `json:"updatename" bson:"updatename"`
}
