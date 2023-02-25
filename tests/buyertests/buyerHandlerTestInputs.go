package buyertests

import (
	"ecomm-alpha/handler"
	"ecomm-alpha/tests/utility"
)

type CreateBuyerAccountTestInput struct {
	Description string
	handler.BuyerSignUpDetails
	utility.RequestConfig
	ExpectedResponseStatusCode int
	ExpectedResponseBody       map[string]interface{}
}

func prepareCreateBuyerAccountTestInputs() *[]CreateBuyerAccountTestInput {

	testInputs := make([]CreateBuyerAccountTestInput, 2)
	testInputs[0].Description = "test with correct inputs"
	testInputs[0].Email = "buyer1@gmail.com"
	testInputs[0].FullName = "dani mason"
	testInputs[0].Password = "!Aa123456"
	testInputs[0].ConfirmPassword = "!Aa123456"
	testInputs[0].RequestRoutePath = "/buyer/signup"
	testInputs[0].RequestMethod = "POST"
	testInputs[0].RequestHeaders = []map[string]string{{"Content-Type": "application/json"}}
	testInputs[0].ExpectedResponseStatusCode = 201
	testInputs[0].ExpectedResponseBody = make(map[string]interface{})
	testInputs[0].ExpectedResponseBody["success"] = false
	testInputs[0].ExpectedResponseBody["message"] = ""
	testInputs[0].ExpectedResponseBody["data"] = ""

	testInputs[1].Description = "test with all empty inputs"
	testInputs[1].Email = ""
	testInputs[1].FullName = ""
	testInputs[1].Password = ""
	testInputs[1].ConfirmPassword = ""
	testInputs[1].RequestRoutePath = "/buyer/signup"
	testInputs[1].RequestMethod = "POST"
	testInputs[1].RequestHeaders = []map[string]string{{"Content-Type": "application/json"}}
	testInputs[1].ExpectedResponseStatusCode = 400
	testInputs[1].ExpectedResponseBody = make(map[string]interface{})
	testInputs[1].ExpectedResponseBody["success"] = false
	testInputs[1].ExpectedResponseBody["message"] = "validation error"
	testInputs[1].ExpectedResponseBody["data"] = []interface{}{"email", "fullName", "password"}

	return &testInputs
}
