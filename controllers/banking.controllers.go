package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/interfaces"
	"main.go/models"
)

type BankingController struct{
	BankingService interfaces.IBank
}

func BankingControllerInit(bankingService interfaces.IBank)BankingController{
	return BankingController{bankingService}
}

func(bc *BankingController)CreateCustomer(ctx *gin.Context){
	var customer *models.Customer

	if err := ctx.ShouldBindJSON(&customer); err!=nil{
		ctx.JSON(http.StatusBadRequest,err.Error())
		return
	}

	newcustomer,err :=bc.BankingService.CreateCustomer(customer)

	if err!=nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"status": "fail","Data":err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated,gin.H{"status":"successfull","Data":newcustomer})


}