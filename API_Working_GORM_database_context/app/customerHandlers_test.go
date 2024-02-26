package app

//"api/service"

//"errors"
//"strconv"

// func Test_should_return_customers_with_status_code_200(t *testing.T){
// 	//Arrange (here we are defining the customerhandler)
//      ctrl := gomock.NewController(t)
// 	 defer ctrl.Finish() //if finish is not called it will not

// 	mockService := service.NewMockCustomerService(ctrl)
// 	dummyCustomers := []dto.CustomerResponse{
// 		{Id: 1002, Name: "Bhavana", City: "Hyderabad", Zipcode: "1023410", Status: "1"},
// 		{Id: 1003, Name: "Chithra", City: "Mysore", Zipcode: "145010", Status: "0"},
//  	}

// 	mockService.EXPECT().GetAllCustomers().Return(dummyCustomers, nil)
// 	ch := CustomerHandlers{mockService}

// 	router := gin.Default()

// 	router.GET("/customers", ch.GetAllCustomer)
// // Here we are creating the http request

// 	req, _ := http.NewRequest(http.MethodGet, "/customers", nil)

// 	// Act
// 	recorder := httptest.NewRecorder()
//router.ServeHTTP(recorder, req)

// 	//Assert
// 	if recorder.Code != http.StatusOK{
// 		t.Error("Failed while testing the status code")
// 	}

// }

// func Test_should_return_the_customer_with_error_message_500(t *testing.T){

// 	//Arrange (here we are defining the customerhandler)
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish() //if finish is not called it will not

//    mockService := service.NewMockCustomerService(ctrl)

//    mockService.EXPECT().GetAllCustomers().Return(nil, errors.New("Unexpected database error"))
//    ch := CustomerHandlers{mockService}
//    router := gin.Default()
//    router.GET("/customers", ch.GetAllCustomer)
// // Here we are creating the http request
//    req, _ := http.NewRequest(http.MethodGet, "/customers", nil)

//    // Act
// 	recorder := httptest.NewRecorder()
// 	router.ServeHTTP(recorder, req)

// 	//Assert
// 	if recorder.Code != http.StatusInternalServerError{
// 		t.Error("Failed while testing the status code")
// 	}
// }

// func Test_Should_return_customer_id_with_the_status_code_200(t *testing.T) {
// 	// Arrange
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	mockService := service.NewMockCustomerService(ctrl)
// 	ch := CustomerHandlers{mockService}
// 	router := gin.Default()

// 	router.GET("/customers/:id", ch.GetCustomerById)

// 	expectedCustomer := &dto.CustomerResponse{Id: 1002, Name: "Bhavana", City: "Hyderabad", Zipcode: "1023410", Status: "1"}
// 	mockService.EXPECT().GetCustomerById(1001).Return(expectedCustomer, nil)

// 	req, _ := http.NewRequest(http.MethodGet, "/customers/1001", nil)

// 	// Act
// 	recorder := httptest.NewRecorder()
// 	router.ServeHTTP(recorder, req)

// 	// Assert
// 	expectedStatusCode := http.StatusOK
// 	if recorder.Code != expectedStatusCode {
// 		t.Errorf("Expected status code %d, but got %d", expectedStatusCode, recorder.Code)
// 	}
// }
