package producttests

import (
	"ecomm-alpha/models"
	"ecomm-alpha/tests/commonData"
	"ecomm-alpha/tests/utility"

	"github.com/google/uuid"
)

type CreateProductTestInput struct {
	Description string
	models.Product
	utility.RequestConfig
	ExpectedResponseStatusCode int
	ExpectedResponseBody       map[string]interface{}
}

func prepareCreateProductTestInputs() *[]CreateProductTestInput {

	testInputs := make([]CreateProductTestInput, 2)

	token := utility.Login(commonData.Seller.Email, commonData.Seller.Password, "/seller/login")

	storeID := "52345971-7564-44f4-a95f-3c5d3ed962d5" //todo: get store id
	testInputs[0].Description = "test with correct inputs"

	testInputs[0].StoreID = uuid.MustParse(storeID)
	testInputs[0].Title = "extension board"
	testInputs[0].Product.Description = "5 plugs"
	testInputs[0].Price = 500
	testInputs[0].Discount = 10
	testInputs[0].AvailableQuantity = 10
	testInputs[0].RequestRoutePath = "/product/"
	testInputs[0].RequestMethod = "POST"

	testInputs[0].RequestHeaders = []map[string]string{{"Content-Type": "application/json", "Authorization": "Bearer " + token}}

	testInputs[0].ExpectedResponseStatusCode = 201

	testInputs[1].Description = "test with all empty inputs"
	testInputs[1].StoreID = uuid.MustParse(storeID)
	testInputs[1].Title = ""
	testInputs[1].Product.Description = ""
	testInputs[1].Price = 0
	testInputs[1].Discount = 0
	testInputs[1].AvailableQuantity = 0
	testInputs[1].RequestRoutePath = "/product/"
	testInputs[1].RequestMethod = "POST"

	testInputs[1].RequestHeaders = []map[string]string{{"Content-Type": "application/json", "Authorization": "Bearer " + token}}
	testInputs[1].ExpectedResponseStatusCode = 400
	testInputs[1].ExpectedResponseBody = make(map[string]interface{})
	testInputs[1].ExpectedResponseBody["success"] = false
	testInputs[1].ExpectedResponseBody["message"] = "validation error"
	testInputs[1].ExpectedResponseBody["data"] = []interface{}{"title", "description", "availableQuantity"}

	return &testInputs
}
