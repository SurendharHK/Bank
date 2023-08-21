package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"main.go/interfaces"
	"main.go/models"
)

type BankingService struct {
	BankingCollection *mongo.Collection
	ctx        context.Context
}

func BankingServiceInit(collection *mongo.Collection, ctx context.Context)interfaces.IBank{
	return &BankingService{collection,ctx}
}

func(b *BankingService)CreateCustomer(customer *models.Customer)(string,error){

	_,err:=b.BankingCollection.InsertOne(b.ctx,customer)

	if err!=nil{
		return "error",err
	}
	return "success",nil



}
