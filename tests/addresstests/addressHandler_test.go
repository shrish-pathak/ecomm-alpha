package addresstests

import (
	"bytes"
	"ecomm-alpha/tests/utility"
	"encoding/json"
	"net/http"
	"testing"
)

func Test_CreateAddress(t *testing.T) {
	testInputs := prepareCreateAddressTestInputs()
	for _, testInput := range *testInputs {
		t.Run(testInput.Description, func(t *testing.T) {
			store := testInput.Address
			storeByte, err := json.Marshal(store)
			if err != nil {
				t.Error(err)
			}
			payload := bytes.NewBuffer(storeByte)
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

			// if caseNum == 0 {
			// 	if testInput.ExpectedResponseBody["success"] != res.(map[string]interface{})["success"].(bool) {
			// 		t.Error("success false")
			// 		t.Error("got: ", res)
			// 	}
			// 	return
			// }

			// if reflect.DeepEqual(testInput.ExpectedResponseBody, res) == false {
			// 	t.Error("wrong response body")
			// 	t.Error("expected: ", testInput.ExpectedResponseBody)
			// 	t.Error("got: ", res)
			// }
		})
	}
}
