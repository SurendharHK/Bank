package interfaces

import "main.go/models"



type IBank interface {
	CreateCustomer(customer *models.Customer)(string,error)
}
