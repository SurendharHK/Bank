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
	ctx := context.TODO()
	mongocollection := mongoClient.Database(constants.ConnectionStrings).Collection("profiles")
	bankingservice := services.BankingServiceInit(mongocollection, ctx)
	bankingController := controllers.BankingControllerInit(bankingservice)
	routes.BankingRoute(server, bankingController)

}

func main() {
	server := gin.Default()
	mongoClient, err := config.ConnectDataBase()
	defer mongoClient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}

	initApp(mongoClient)
	fmt.Println("server is running on", constants.Port)
	log.Fatal(server.Run(constants.Port))

}
