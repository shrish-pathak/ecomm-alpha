package producttests

import (
	"bytes"
	"ecomm-alpha/tests/utility"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func Test_CreateProduct(t *testing.T) {
	testInputs := prepareCreateProductTestInputs()

	for caseNum, testInput := range *testInputs {
		t.Run(testInput.Description, func(t *testing.T) {
			product := testInput.Product
			productByte, err := json.Marshal(product)
			if err != nil {
				t.Error(err)
			}

			payload := bytes.NewBuffer(productByte)
			req, err := http.NewRequest(testInput.RequestMethod, utility.BaseUrl+testInput.RequestRoutePath, payload)

			for _, header := range testInput.RequestHeaders {
				for k, v := range header {
					req.Header.Set(k, v)
				}
			}

			if err != nil {
				t.Error(err)
			}

			res, statusCode, err := utility.GetResponse(req)

			if testInput.ExpectedResponseStatusCode != statusCode {
				t.Error("wrong status code")
				t.Error("expected: ", testInput.ExpectedResponseStatusCode)
				t.Error("got: ", statusCode)
				return
			}

			if caseNum == 0 {
				return
			}

			if reflect.DeepEqual(testInput.ExpectedResponseBody, res) == false {
				t.Error("wrong response body")
				t.Error("expected: ", testInput.ExpectedResponseBody)
				t.Error("got: ", res)
			}
		})
	}
}
