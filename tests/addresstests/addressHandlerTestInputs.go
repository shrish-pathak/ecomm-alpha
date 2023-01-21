package addresstests

import (
	"ecomm-alpha/models"
	"ecomm-alpha/tests/commonData"
	"ecomm-alpha/tests/utility"
)

type CreateAddressTestInput struct {
	Description string
	models.Address
	utility.RequestConfig
	ExpectedResponseStatusCode int
	ExpectedResponseBody       map[string]interface{}
}

func prepareCreateAddressTestInputs() *[]CreateAddressTestInput {

	token := utility.Login(commonData.Seller.Email, commonData.Seller.Password)

	testInputs := make([]CreateAddressTestInput, 1)
	testInputs[0].Description = "test with correct inputs"
	testInputs[0].MobileNo = "9988776655"
	testInputs[0].City = "london"
	testInputs[0].State = "london"
	testInputs[0].Zip = "EC1A"
	testInputs[0].Country = "United Kingdom"
	testInputs[0].Address.Address = "City of Westminster, London, SW1"

	testInputs[0].RequestRoutePath = "/address/"
	testInputs[0].RequestMethod = "POST"
	testInputs[0].RequestHeaders = []map[string]string{{"Content-Type": "application/json", "Authorization": "Bearer " + token}}
	testInputs[0].ExpectedResponseStatusCode = 201
	// testInputs[0].ExpectedResponseBody = make(map[string]interface{})
	// testInputs[0].ExpectedResponseBody["success"] = true
	// testInputs[0].ExpectedResponseBody["message"] = ""
	// testInputs[0].ExpectedResponseBody["data"] = nil

	return &testInputs
}
