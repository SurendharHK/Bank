package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"main.go/interfaces"
	"main.go/models"
)

type BankingController struct {
	TransactionService interfaces.IBank
}

func InitBankingController(profileService interfaces.IBank) BankingController {
	return BankingController{profileService}
}

func (pc *BankingController) CreateTransaction(ctx *gin.Context) {
	var profile *models.Customer
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newProfile, err := pc.TransactionService.CreateCustomer(profile)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newProfile})
}

func(pc *BankingController)GetCustomers(ctx *gin.Context){
	


	customers,err:=pc.TransactionService.GetCustomers()
	if err!=nil{
		ctx.JSON(http.StatusNotFound,gin.H{"status":"fail","message":err.Error()})
	}
	ctx.JSON(http.StatusFound,gin.H{"status":"success","message":customers})

}

