package utility

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var client *http.Client
var BaseUrl = "http://localhost:5000/api/v1"
var showtLogs = false
var initCallCounter = 0

func init() {
	initCallCounter++
	log.Println("initCallCounter", initCallCounter)
	client = new(http.Client)
}

type RequestConfig struct {
	RequestRoutePath string
	RequestMethod    string
	RequestHeaders   []map[string]string
}

func GetResponse(req *http.Request) (interface{}, int, error) {

	var resp *http.Response
	var err error

	resp, err = client.Do(req)

	if err != nil {
		return nil, resp.StatusCode, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, resp.StatusCode, err
	}

	if resp.StatusCode == 204 {
		return nil, resp.StatusCode, nil
	}
	var data interface{}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, resp.StatusCode, err
	}
	return data, resp.StatusCode, nil
}
