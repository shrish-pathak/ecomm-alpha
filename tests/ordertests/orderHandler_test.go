package ordertests

import (
	"bytes"
	"ecomm-alpha/tests/utility"
	"encoding/json"
	"net/http"
	"testing"
)

func Test_PlaceOrder(t *testing.T) {
	testInputs := preparePlaceOrderTestInputs()

	for _, testInput := range *testInputs {
		t.Run(testInput.Description, func(t *testing.T) {
			order := testInput.Order
			orderByte, err := json.Marshal(order)

			if err != nil {
				t.Error(err)
			}

			payload := bytes.NewBuffer(orderByte)

			req, err := http.NewRequest(testInput.RequestMethod, utility.BaseUrl+testInput.RequestRoutePath, payload)
			for _, header := range testInput.RequestHeaders {
				for k, v := range header {
					req.Header.Set(k, v)
				}
			}
			if err != nil {
				t.Error(err)
			}

			_, statusCode, err := utility.GetResponse(req)

			if testInput.ExpectedResponseStatusCode != statusCode {
				t.Error("wrong status code")
				t.Error("expected: ", testInput.ExpectedResponseStatusCode)
				t.Error("got: ", statusCode)
				return
			}

		})
	}
}
