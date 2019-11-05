package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type ResponseResult struct {
	Body       []byte
	StatusCode int
}

func Get(url string) (*ResponseResult, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	return &ResponseResult{Body: data, StatusCode: response.StatusCode}, err
}

func Request(url, method string, body interface{}, headers ...map[string]string) (*ResponseResult, error) {
	jsonBody, _ := json.Marshal(body)
	request, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-type", "application/json")
	if len(headers) > 0 {
		for k, v := range headers[0] {
			request.Header.Set(k, v)
		}
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	return &ResponseResult{Body: data, StatusCode: response.StatusCode}, err
}
