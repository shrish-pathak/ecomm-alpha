package sellertests

import (
	"ecomm-alpha/models"
	"ecomm-alpha/tests/utility"
)

type CreateStoreTestInput struct {
	Description string
	models.Store
	utility.RequestConfig
	ExpectedResponseStatusCode int
	ExpectedResponseBody       map[string]interface{}
}

func prepareCreateStoreTestInputs() *[]CreateStoreTestInput {

	testInputs := make([]CreateStoreTestInput, 1)
	testInputs[0].Description = "test with correct inputs"
	testInputs[0].Name = "potion store"
	testInputs[0].Store.Description = "potions"
	testInputs[0].RequestRoutePath = "/store/"
	testInputs[0].RequestMethod = "POST"
	testInputs[0].RequestHeaders = []map[string]string{{"Content-Type": "application/json"}}
	testInputs[0].ExpectedResponseStatusCode = 201
	testInputs[0].ExpectedResponseBody = make(map[string]interface{})
	testInputs[0].ExpectedResponseBody["success"] = false
	testInputs[0].ExpectedResponseBody["message"] = ""
	testInputs[0].ExpectedResponseBody["data"] = ""

	testInputs[1].Description = "test with empty name input"
	testInputs[1].Name = ""
	testInputs[1].Store.Description = "potions"
	testInputs[1].RequestRoutePath = "/store/"
	testInputs[1].RequestMethod = "POST"
	testInputs[1].RequestHeaders = []map[string]string{{"Content-Type": "application/json"}}
	testInputs[1].ExpectedResponseStatusCode = 400
	testInputs[1].ExpectedResponseBody = make(map[string]interface{})
	testInputs[1].ExpectedResponseBody["success"] = false
	testInputs[1].ExpectedResponseBody["message"] = "validation error"
	testInputs[1].ExpectedResponseBody["data"] = []interface{}{"name"}

	testInputs[2].Description = "test with empty description input"
	testInputs[2].Name = "potions"
	testInputs[2].Store.Description = ""
	testInputs[2].RequestRoutePath = "/store/"
	testInputs[2].RequestMethod = "POST"
	testInputs[2].RequestHeaders = []map[string]string{{"Content-Type": "application/json"}}
	testInputs[2].ExpectedResponseStatusCode = 400
	testInputs[2].ExpectedResponseBody = make(map[string]interface{})
	testInputs[2].ExpectedResponseBody["success"] = false
	testInputs[2].ExpectedResponseBody["message"] = "validation error"
	testInputs[2].ExpectedResponseBody["data"] = []interface{}{"description"}

	testInputs[3].Description = "test with correct payload and wrong content-type"
	testInputs[3].Name = "potion store"
	testInputs[3].Store.Description = "potions"
	testInputs[3].RequestRoutePath = "/store/"
	testInputs[3].RequestMethod = "POST"
	testInputs[3].RequestHeaders = []map[string]string{{"Content-Type": "application/text"}}
	testInputs[3].ExpectedResponseStatusCode = 422
	testInputs[3].ExpectedResponseBody = make(map[string]interface{})
	testInputs[3].ExpectedResponseBody["success"] = false
	testInputs[3].ExpectedResponseBody["message"] = "Unprocessable Entity"
	testInputs[3].ExpectedResponseBody["data"] = nil

	return &testInputs
}
