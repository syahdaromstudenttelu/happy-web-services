package helper

import (
	"encoding/json"
	"happy-admin-service/model/web"
	"io"
	"net/http"
)

func CreateHttpRequestService(method string, httpUrl string, body io.Reader) web.ServiceResponse {
	request, err := http.NewRequest(method, httpUrl, body)
	DoPanicIfError(err)

	client := &http.Client{}
	response, err := client.Do(request)
	DoPanicIfError(err)
	defer response.Body.Close()

	serviceResponse := web.ServiceResponse{}
	err = json.NewDecoder(response.Body).Decode(&serviceResponse)
	DoPanicIfError(err)

	return serviceResponse
}
