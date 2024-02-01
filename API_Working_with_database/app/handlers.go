package app

import (
	"api/service"
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
	id := c.Param("id")
	idint, _ := strconv.Atoi(id)

	customer, _ := ch.service.GetCustomerById(idint)
	c.JSON(http.StatusOK, customer)
}
