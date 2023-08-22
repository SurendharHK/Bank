package services

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main.go/interfaces"
	"main.go/models"
)

type BankingService struct {
	BankingCollection *mongo.Collection
	ctx               context.Context
}



// GetCustomers implements interfaces.IBank.
func (t *BankingService) GetCustomers() ([]*models.Customer, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{}}
	result, err := t.BankingCollection.Find(ctx, filter)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		//do something
		fmt.Println(result)
		//build the array of products for the cursor that we received.
		var customers []*models.Customer
		for result.Next(ctx) {
			product := &models.Customer{}
			err := result.Decode(product)

			if err != nil {
				return nil, err
			}
			//fmt.Println(product)
			customers = append(customers, product)
		}
		if err := result.Err(); err != nil {
			return nil, err
		}
		if len(customers) == 0 {
			return []*models.Customer{}, nil
		}

		return customers, nil
	}
}

func NewBankingServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.IBank {
	return &BankingService{collection, ctx}

}

// CreateTransaction implements interfaces.ITransaction.
func (t *BankingService) CreateCustomer(customer *models.Customer) (string, error) {

	//hashedPassword, _ := utils.HashPassword(user.Password)

	_, err := t.BankingCollection.InsertOne(t.ctx, &customer)
	if err != nil {
		return "error", nil
	}

	return "success", nil
}
