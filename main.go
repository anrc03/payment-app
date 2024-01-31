package main

import (
	"enigma.com/m/controllers"
	"enigma.com/m/initializers"
	"enigma.com/m/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	//auth
	r.POST("/api/auth/signup", controllers.SignUp)
	r.POST("/api/auth/login", controllers.Login)
	//customer
	r.GET("/api/customer", middleware.RequireAuth, controllers.GetAllCustomer)
	r.POST("/api/auth/register", controllers.RegisterCustomer)
	//merchant
	r.GET("/api/merchant", middleware.RequireAuth, controllers.GetAllMerchant)
	r.POST("/api/auth/register/merchant-store", controllers.RegisterMerchant)
	//store
	r.POST("/api/auth/register/store", controllers.RegisterStore)
	r.GET("/api/store", middleware.RequireAuth, controllers.GetAllStore)
	//transaction
	r.POST("/api/pay", middleware.RequireAuth, controllers.CreateTransaction)
	r.GET("/api/pay", middleware.RequireAuth, controllers.GetTransactionHistory)

	r.Run()
}
