package domain

// Here we are creating the the stub adapter(ie we had to collect the mock adapter but here we are not )
type CustomerRepositoryStub struct {
	customers []Customer
}

// here we are creating a reciver function and attaching the struct
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// helper function , this func will be creating a new customer repository stub

// func NewCustomerRepositoryStub() CustomerRepositoryStub {
// 	customers := []Customer{
// 		{Id: "1001", Name: "Akash", City: "Bangalore", Zipcode: "100010", DateofBirth: "01/01/1998", Status: "1"},
// 		{Id: "1002", Name: "Bhavana", City: "Hyderabad", Zipcode: "1023410", DateofBirth: "05/07/2001", Status: "1"},
// 		{Id: "1003", Name: "Chithra", City: "Mysore", Zipcode: "145010", DateofBirth: "11/09/1995", Status: "0"},
// 	}
// 	return CustomerRepositoryStub{customers: customers}
// }

func (c CustomerRepositoryStub) GetCustomerById(int) (Customer, error) {
	return Customer{}, nil
}
