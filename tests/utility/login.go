package utility

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func login(email, password string) string {
	client := new(http.Client)

	loginData := map[string]string{
		"email":    email,
		"password": password,
	}
	LDBytes, err := json.Marshal(loginData)
	if err != nil {
		log.Println(err)
	}
	payload := bytes.NewBuffer(LDBytes)
	req, err := http.NewRequest("POST", BaseUrl+"/seller/login", payload)
	if err != nil {
		log.Println(err)
	}

	res, err := client.Do(req)

	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	rByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(rByte, &data)
	if err != nil {
		log.Println(err)
	}

	return data["data"].(string)
}
