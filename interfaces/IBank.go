package interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
	"main.go/models"
)

type IBank interface {
	CreateCustomer(customer *models.Customer) (string, error)
	GetCustomers() ([]*models.Customer, error)
	UpdateCustomer(intialName string, updateName string) (*mongo.UpdateResult, error)
	DeleteCustomer(name string)(*mongo.DeleteResult,error)
}
