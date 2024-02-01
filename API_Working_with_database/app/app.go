package app

import (
	"api/domain" // Assuming CustomerHandlers is defined in handlers package
	"api/service"
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	// Complete wiring from rest handlers to adapter
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Defining the routes
	router.GET("/customers", ch.GetAllCustomer)
	router.GET("/customers/:id", ch.GetCustomerById)

	err := router.Run("localhost:8000")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
