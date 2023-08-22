package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
)

func BankingRoute(route *gin.Engine,controller controllers.BankingController){
	route.POST("/banking/create",controller.CreateTransaction)
	route.GET("/customers",controller.GetCustomers)
}