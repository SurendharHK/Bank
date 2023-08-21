package models

type Customer struct {
	Name     string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
}
