package app

import (
	"api/domain"
	"api/service"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

// This will be having the dependencies of the service

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomer(c *gin.Context) {
	// customers := []Customer{
	// 	{Name: "Kavya", City: "Banglore", Zipcode: "560003"},
	// 	{Name: "Prakurthi", City: "Hyderabad", Zipcode: "185754"},
	// 	{Name: "Tapasya", City: "Banglore", Zipcode: "560043"},
	// }

	customers, _ := ch.service.GetAllCustomers()

	if c.GetHeader("Content-type") == "application/xml" {
		c.XML(http.StatusOK, customers)
	} else {
		c.JSON(http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) GetCustomerById(c *gin.Context) {
	id := c.Query("id") // Now the customer is accepting values in the query format
	idint, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer id the customer id , must be numeric..."})
		return
	}

	customer, err := ch.service.GetCustomerById(idint)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("%s", err.Error())})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (ch *CustomerHandlers) CreateCustomer(c *gin.Context) {
	var newCustomer domain.Customer

	if err := c.BindJSON(&newCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	_, err := ch.service.CreateCustomer(newCustomer) // Corrected method call
	if err != nil {
		log.Println("Error creating customer:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": "Customer created successfully"})

}
