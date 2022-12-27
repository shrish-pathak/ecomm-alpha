package sellertests

import (
	"ecomm-alpha/handler"
	"ecomm-alpha/tests/utility"
)

type CreateSellerAccountTestInput struct {
	Description string
	handler.SellerSignUpDetails
	utility.RequestConfig
	ExpectedResponseStatusCode int
	ExpectedResponseBody       map[string]interface{}
}

func prepareCreateSellerAccountTestInputs() *[]CreateSellerAccountTestInput {

	testInputs := make([]CreateSellerAccountTestInput, 9)
	testInputs[0].Description = "test with correct inputs"
	testInputs[0].Email = "harry@gmail.com"
	testInputs[0].FullName = "harry potter"
	testInputs[0].Password = "expecto patronum"
	testInputs[0].ConfirmPassword = "expecto patronum"
	testInputs[0].RequestRoutePath = "/seller/signup"
	testInputs[0].RequestMethod = "POST"
	testInputs[0].RequestHeaders = []map[string]string{{"Content-Type": "application/json"}}
	testInputs[0].ExpectedResponseStatusCode = 201
	testInputs[0].ExpectedResponseBody = make(map[string]interface{})
	testInputs[0].ExpectedResponseBody["success"] = false
	testInputs[0].ExpectedResponseBody["message"] = ""
	testInputs[0].ExpectedResponseBody["data"] = ""

	testInputs[1].Description = "test with invalid email input"
	testInputs[1].Email = "harrygmail.com"
	testInputs[1].FullName = "harry potter"
	testInputs[1].Password = "expecto patronum"
	testInputs[1].ConfirmPassword = "expecto patronum"
	testInputs[1].RequestRoutePath = "/seller/signup"
	testInputs[1].RequestMethod = "POST"
	testInputs[1].RequestHeaders = []map[string]string{{"Content-Type": "application/json"}}
	testInputs[1].ExpectedResponseStatusCode = 400
	testInputs[1].ExpectedResponseBody = make(map[string]interface{})
	testInputs[1].ExpectedResponseBody["success"] = false
	testInputs[1].ExpectedResponseBody["message"] = "validation error"
	testInputs[1].ExpectedResponseBody["data"] = []interface{}{"email"}

	testInputs[2].Description = "test with unmatched ConfirmPassword input"
	testInputs[2].Email = "harry@gmail.com"
	testInputs[2].FullName = "harry potter"
	testInputs[2].Password = "expecto patronum"
	testInputs[2].ConfirmPassword = "expelliarmus"
	testInputs[2].RequestRoutePath = "/seller/signup"
	testInputs[2].RequestMethod = "POST"
	testInputs[2].RequestHeaders = []map[string]string{{"Content-Type": "application/json"}}
	testInputs[2].ExpectedResponseStatusCode = 400
	testInputs[2].ExpectedResponseBody = make(map[string]interface{})
	testInputs[2].ExpectedResponseBody["success"] = false
	testInputs[2].ExpectedResponseBody["message"] = "validation error"
	testInputs[2].ExpectedResponseBody["data"] = []interface{}{"confirmPassword"}

	testInputs[3].Description = "test with empty email"
	testInputs[3].Email = ""
	testInputs[3].FullName = "harry potter"
	testInputs[3].Password = "expecto patronum"
	testInputs[3].ConfirmPassword = "expecto patronum"
	testInputs[3].RequestRoutePath = "/seller/signup"
	testInputs[3].RequestMethod = "POST"
	testInputs[3].RequestHeaders = []map[string]string{{"Content-Type": "application/json"}}
	testInputs[3].ExpectedResponseStatusCode = 400
	testInputs[3].ExpectedResponseBody = make(map[string]interface{})
	testInputs[3].ExpectedResponseBody["success"] = false
	testInputs[3].ExpectedResponseBody["message"] = "validation error"
	testInputs[3].ExpectedResponseBody["data"] = []interface{}{"email"}

	testInputs[4].Description = "test with all fields missing"
	testInputs[4].RequestRoutePath = "/seller/signup"
	testInputs[4].RequestMethod = "POST"
	testInputs[4].RequestHeaders = []map[string]string{{"Content-Type": "application/json"}}
	testInputs[4].ExpectedResponseStatusCode = 400
	testInputs[4].ExpectedResponseBody = make(map[string]interface{})
	testInputs[4].ExpectedResponseBody["success"] = false
	testInputs[4].ExpectedResponseBody["message"] = "validation error"
	testInputs[4].ExpectedResponseBody["data"] = []interface{}{"email", "fullName", "password"}

	testInputs[5].Description = "test with empty full name"
	testInputs[5].Email = "harry@gmail.com"
	testInputs[5].FullName = ""
	testInputs[5].Password = "expecto patronum"
	testInputs[5].ConfirmPassword = "expecto patronum"
	testInputs[5].RequestRoutePath = "/seller/signup"
	testInputs[5].RequestMethod = "POST"
	testInputs[5].RequestHeaders = []map[string]string{{"Content-Type": "application/json"}}
	testInputs[5].ExpectedResponseStatusCode = 400
	testInputs[5].ExpectedResponseBody = make(map[string]interface{})
	testInputs[5].ExpectedResponseBody["success"] = false
	testInputs[5].ExpectedResponseBody["message"] = "validation error"
	testInputs[5].ExpectedResponseBody["data"] = []interface{}{"fullName"}

	testInputs[6].Description = "test with empty password"
	testInputs[6].Email = "harry@gmail.com"
	testInputs[6].FullName = "harry potter"
	testInputs[6].Password = ""
	testInputs[6].ConfirmPassword = ""
	testInputs[6].RequestRoutePath = "/seller/signup"
	testInputs[6].RequestMethod = "POST"
	testInputs[6].RequestHeaders = []map[string]string{{"Content-Type": "application/json"}}
	testInputs[6].ExpectedResponseStatusCode = 400
	testInputs[6].ExpectedResponseBody = make(map[string]interface{})
	testInputs[6].ExpectedResponseBody["success"] = false
	testInputs[6].ExpectedResponseBody["message"] = "validation error"
	testInputs[6].ExpectedResponseBody["data"] = []interface{}{"password"}

	testInputs[7].Description = "test with correct payload and wrong content-type"
	testInputs[7].Email = "harry@gmail.com"
	testInputs[7].FullName = "harry potter"
	testInputs[7].Password = ""
	testInputs[7].ConfirmPassword = ""
	testInputs[7].RequestRoutePath = "/seller/signup"
	testInputs[7].RequestMethod = "POST"
	testInputs[7].RequestHeaders = []map[string]string{{"Content-Type": "application/text"}}
	testInputs[7].ExpectedResponseStatusCode = 422
	testInputs[7].ExpectedResponseBody = make(map[string]interface{})
	testInputs[7].ExpectedResponseBody["success"] = false
	testInputs[7].ExpectedResponseBody["message"] = "Unprocessable Entity"
	testInputs[7].ExpectedResponseBody["data"] = nil

	testInputs[8].Description = "test with already existing user"
	testInputs[8].Email = "harry@gmail.com"
	testInputs[8].FullName = "harry potter"
	testInputs[8].Password = "expecto patronum"
	testInputs[8].ConfirmPassword = "expecto patronum"
	testInputs[8].RequestRoutePath = "/seller/signup"
	testInputs[8].RequestMethod = "POST"
	testInputs[8].RequestHeaders = []map[string]string{{"Content-Type": "application/json"}}
	testInputs[8].ExpectedResponseStatusCode = 400
	testInputs[8].ExpectedResponseBody = make(map[string]interface{})
	testInputs[8].ExpectedResponseBody["success"] = false
	testInputs[8].ExpectedResponseBody["message"] = "user already exists"
	testInputs[8].ExpectedResponseBody["data"] = nil

	return &testInputs
}
