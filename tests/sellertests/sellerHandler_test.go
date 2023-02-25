package sellertests

import (
	"bytes"
	"ecomm-alpha/config"
	"ecomm-alpha/database"
	"ecomm-alpha/models"
	"ecomm-alpha/tests/utility"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"

	"github.com/golang-jwt/jwt/v4"
)

func Test_CreateSellerAccount(t *testing.T) {
	testInputs := prepareCreateSellerAccountTestInputs()
	for caseNum, testInput := range *testInputs {
		t.Run(testInput.Description, func(t *testing.T) {
			sellerSignUpDetails := testInput.SellerSignUpDetails
			sellerSUDByte, err := json.Marshal(sellerSignUpDetails)
			if err != nil {
				t.Error(err)
			}
			payload := bytes.NewBuffer(sellerSUDByte)
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
				tokenString := res.(map[string]interface{})["data"].(string)

				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}

					return []byte(config.Config("SECRET")), nil
				})

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					log.Println(claims)
				} else {
					t.Error("token validation error: ", err)
				}
				return
			}
			if reflect.DeepEqual(testInput.ExpectedResponseBody, res) == false {
				t.Error("wrong response body")
				t.Error("expected: ", testInput.ExpectedResponseBody)
				t.Error("got: ", res)
			}
		})
	}

	database.ConnectDB()
	db := database.DB
	testUser := new(models.Seller)
	err := db.Raw("delete from sellers where email = ? returning *;", (*testInputs)[0].Email).Scan(&testUser).Error

	if testUser == nil && err != nil {
		t.Error("failed to delete test user signup details")
	}

}
