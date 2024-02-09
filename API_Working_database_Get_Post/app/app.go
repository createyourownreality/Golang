package app

import (
	"api/domain" // Assuming CustomerHandlers is defined in handlers package
	"api/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// func sanityCheck() {
// 	addr := os.Getenv("ADDR")
// 	port := os.Getenv("PORT")
// 	user := os.Getenv("DB_USER")
// 	password := os.Getenv("DB_PASSWORD")
// 	dbname := os.Getenv("DB_NAME")
// 	sslmode := os.Getenv("DB_SSLMODE")

// 	if addr == "" || port == "" || user == "" || password == "" || dbname == "" || sslmode == "" {
// 		log.Fatal("Environment variables are not defined...")
// 	}
// }

func Start() {

	//sanityCheck()

	router := gin.Default()

	// Complete wiring from rest handlers to adapter
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Defining the routes
	router.GET("/customers", ch.GetAllCustomer)
	router.GET("/customers/:id", ch.GetCustomerById)
	router.POST("/customers", ch.CreateCustomer)

	addr := os.Getenv("ADDR")
	port := os.Getenv("PORT")
	if addr == "" || port == "" {
		log.Fatal("ADDR and PORT environment variables are not set")
	}

	err := router.Run(addr + ":" + port)
	if err != nil {
		log.Fatal(err)
	}

}
