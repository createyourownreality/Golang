package app

import (
	"api/domain"
	"api/logger"
	"api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// This will be having the dependencies of the service

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomer(c *gin.Context) {
	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve customers"})
		return
	}

	// Respond with JSON
	c.JSON(http.StatusOK, customers)
}

func (ch *CustomerHandlers) GetCustomerById(c *gin.Context) {
	//id := c.Query("id") // Now the customer is accepting values in the query format
	
    customer_id := c.Param("id")
 
	idint, err := strconv.Atoi(customer_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid customer id the customer id , must be numeric..."})
		return
	}


	customer, err := ch.service.GetCustomerById(idint)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found...."})
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
		logger.Error("Error creating customer:")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": "Customer created successfully"})

}
