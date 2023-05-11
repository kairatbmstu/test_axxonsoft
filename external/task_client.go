package external

import (
	"io/ioutil"
	"log"
	"net/http"
)

type TaskClient struct {
}

type HttpRequest struct {
	Method         string
	Url            string
	RequestHeaders map[string]string
	RequestBody    string
}

type HttpResponse struct {
	Method          string
	Url             string
	Status          int
	ResponseHeaders map[string]string
	ResponseBody    string
}

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
	log.Println("httpResponse: ", httpResponse)
	return &httpResponse, nil
}
