package external

import (
	"io/ioutil"
	"log"
	"net/http"
)

/*
The TaskClient type represents an HTTP client for making requests.
It does not have any specific properties or methods but is used as a receiver for the DoHttpRequest method.
*/
type TaskClient struct {
}

/**
The HttpRequest type represents an HTTP request. It contains the following fields:

Method: A string representing the HTTP method for the request (e.g., "GET", "POST", "PUT", etc.).
Url: A string representing the URL to which the request should be sent.
RequestHeaders: A map of string key-value pairs representing the request headers.
RequestBody: A string representing the body of the request.
*/
type HttpRequest struct {
	Method         string
	Url            string
	RequestHeaders map[string]string
	RequestBody    string
}

/**
The HttpResponse type represents an HTTP response. It contains the following fields:

Method: A string representing the HTTP method used for the corresponding request.
Url: A string representing the URL to which the request was sent.
Status: An integer representing the HTTP status code of the response.
ResponseHeaders: A map of string key-value pairs representing the response headers.
ResponseBody: A string representing the body of the response.
*/
type HttpResponse struct {
	Method          string
	Url             string
	Status          int
	ResponseHeaders map[string]string
	ResponseBody    string
}

/**
The DoHttpRequest method is used to send an HTTP request and retrieve the corresponding response.
It takes an HttpRequest object as a parameter and returns an HttpResponse object and an error (if any).
*/
func (t *TaskClient) DoHttpRequest(HttpRequest *HttpRequest) (*HttpResponse, error) {
	var httpResponse = HttpResponse{}
	request, err := http.NewRequest(HttpRequest.Method, HttpRequest.Url, nil)
	if err != nil {
		log.Println("client: could not create request: ", err)
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println("client: error making http request: ", err)
		return nil, err
	}

	httpResponse.Status = response.StatusCode

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("client: could not read response body: ", err)
		return nil, err
	}
	httpResponse.ResponseBody = string(responseBody)
	//log.Println("httpResponse: ", httpResponse)
	return &httpResponse, nil
}
