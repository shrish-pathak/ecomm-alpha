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

	token := utility.Login(commonData.Seller.Email, commonData.Seller.Password, "/seller/login")

	testInputs := make([]CreateAddressTestInput, 2)
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

	testInputs[1].Description = "test with all empty inputs"
	testInputs[1].MobileNo = ""
	testInputs[1].City = ""
	testInputs[1].State = ""
	testInputs[1].Zip = ""
	testInputs[1].Country = ""
	testInputs[1].Address.Address = ""

	testInputs[1].RequestRoutePath = "/address/"
	testInputs[1].RequestMethod = "POST"
	testInputs[1].RequestHeaders = []map[string]string{{"Content-Type": "application/json", "Authorization": "Bearer " + token}}
	testInputs[1].ExpectedResponseStatusCode = 400
	testInputs[1].ExpectedResponseBody = make(map[string]interface{})
	testInputs[1].ExpectedResponseBody["success"] = false
	testInputs[1].ExpectedResponseBody["message"] = "validation error"
	testInputs[1].ExpectedResponseBody["data"] = []string{"mobule num", "city", "state", "zip", "country", "address"}

	return &testInputs
}
