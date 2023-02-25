package ordertests

import (
	"ecomm-alpha/models"
	"ecomm-alpha/tests/commonData"
	"ecomm-alpha/tests/utility"

	"github.com/google/uuid"
)

type PlaceOrderTestInput struct {
	Description string
	models.Order
	utility.RequestConfig
	ExpectedResponseStatusCode int
	ExpectedResponseBody       map[string]interface{}
}

func preparePlaceOrderTestInputs() *[]PlaceOrderTestInput {
	testInputs := make([]PlaceOrderTestInput, 1)

	token := utility.Login(commonData.Buyer.Email, commonData.Buyer.Password, "/buyer/login")

	addressID := "df3ebf12-e967-4ecb-8833-fe47ab951943" //Todo:get buyer addressID

	testInputs[0].Description = "test with correct inputs"
	testInputs[0].AddressID = uuid.MustParse(addressID)
	testInputs[0].RequestRoutePath = "/order/"
	testInputs[0].RequestMethod = "POST"
	testInputs[0].RequestHeaders = []map[string]string{{"Content-Type": "application/json", "Authorization": "Bearer " + token}}
	testInputs[0].ExpectedResponseStatusCode = 201
	testInputs[0].ExpectedResponseBody = make(map[string]interface{})
	testInputs[0].ExpectedResponseBody["success"] = true
	testInputs[0].ExpectedResponseBody["message"] = ""
	testInputs[0].ExpectedResponseBody["data"] = ""

	return &testInputs

}
