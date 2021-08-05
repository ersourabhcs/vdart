package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var client = &http.Client{}

type ResponseHandler struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func WebResponse(data interface{}, message string, status bool) []byte {
	jsonString := ResponseHandler{
		status,
		message,
		data,
	}
	result, err := json.MarshalIndent(jsonString, "", "    ")
	if err != nil {
		log.Println(err)
	}
	return result
}

// ErrorResponse is a wrapper function for returning a web response that includes a standard text message.
func ErrorResponse(w http.ResponseWriter, r *http.Request, code int, data interface{}, message string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(WebResponse(data, message, false))
}

// ErrorResponse is a wrapper function for returning a web response that includes a standard text message.
func SuccessResponse(w http.ResponseWriter, r *http.Request, code int, data interface{}, message string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(WebResponse(data, message, true))
}

// RequestAPIData:  Calls a (API) URL and return the data from the request.
func RequestAPIData(method, url, postdata string, headers map[string]string) ([]byte, int, error) {

	req, err := http.NewRequest(method, url, strings.NewReader(postdata))
	if err != nil {
		return nil, 500, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 500, err
	}
	statusCode := resp.StatusCode

	// read body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, statusCode, err
	}
	return body, statusCode, nil
}
