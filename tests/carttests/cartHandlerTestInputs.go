package carttests

import (
	"ecomm-alpha/models"
	"ecomm-alpha/tests/commonData"
	"ecomm-alpha/tests/utility"
	"math"

	"github.com/google/uuid"
)

type CartItemTestInput struct {
	Description string
	models.CartItem
	utility.RequestConfig
	ExpectedResponseStatusCode int
	ExpectedResponseBody       map[string]interface{}
}

func prepareAddToCartTestInputs() *[]CartItemTestInput {
	testInputs := make([]CartItemTestInput, 3)

	token := utility.Login(commonData.Buyer.Email, commonData.Buyer.Password, "/buyer/login")

	productID := "1d0fe3fa-acd0-4295-be28-ce0459343b59" //TODO: get product id
	testInputs[0].Description = "test with correct inputs"
	testInputs[0].ProductID = uuid.MustParse(productID)
	testInputs[0].Quantity = 1
	testInputs[0].RequestMethod = "POST"
	testInputs[0].RequestRoutePath = "/cart/"
	testInputs[0].RequestHeaders = []map[string]string{{"Content-Type": "application/json", "Authorization": "Bearer " + token}}

	testInputs[0].ExpectedResponseStatusCode = 201

	testInputs[1].Description = "item out of stock test"
	testInputs[1].ProductID = uuid.MustParse(productID)
	testInputs[1].Quantity = math.MaxInt
	testInputs[1].RequestMethod = "POST"
	testInputs[1].RequestRoutePath = "/cart/"
	testInputs[1].RequestHeaders = []map[string]string{{"Content-Type": "application/json", "Authorization": "Bearer " + token}}
	testInputs[1].ExpectedResponseStatusCode = 417
	testInputs[1].ExpectedResponseBody = make(map[string]interface{})
	testInputs[1].ExpectedResponseBody["success"] = false
	testInputs[1].ExpectedResponseBody["message"] = "item out of stock"
	testInputs[1].ExpectedResponseBody["data"] = nil

	testInputs[2].Description = "invalid quantity test"
	testInputs[2].ProductID = uuid.MustParse(productID)
	testInputs[2].Quantity = 0
	testInputs[2].RequestMethod = "POST"
	testInputs[2].RequestRoutePath = "/cart/"
	testInputs[2].RequestHeaders = []map[string]string{{"Content-Type": "application/json", "Authorization": "Bearer " + token}}
	testInputs[2].ExpectedResponseStatusCode = 400
	testInputs[2].ExpectedResponseBody = make(map[string]interface{})
	testInputs[2].ExpectedResponseBody["success"] = false
	testInputs[2].ExpectedResponseBody["message"] = "invalid quantity"
	testInputs[2].ExpectedResponseBody["data"] = nil

	return &testInputs
}
