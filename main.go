package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"main.go/config"
	"main.go/constants"
	"main.go/controllers"
	"main.go/routes"
	"main.go/services"
)

var (
	mongoClient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)


func initApp(mongoClient *mongo.Client) {
	ctx = context.TODO()
	transactionCollection := mongoClient.Database(constants.DatabaseName).Collection("bank")
	transactionService := services.NewBankingServiceInit(transactionCollection, ctx)
	transactionController := controllers.InitBankingController(transactionService)
	routes.BankingRoute(server, transactionController)
}

func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	
	initApp(mongoclient)
	fmt.Println("server running on port", constants.Port)
	log.Fatal(server.Run(constants.Port))
}
