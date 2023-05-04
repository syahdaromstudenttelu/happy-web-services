package helper

import (
	"encoding/json"
	"happy-product-service/exception"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result any) {
	jsonDecoder := json.NewDecoder(request.Body)
	err := jsonDecoder.Decode(result)

	if err != nil {
		panic(exception.NewReqBodyMalformedError(err.Error()))
	}
}

func WriteToResponseBody(writer http.ResponseWriter, response any) {
	writer.Header().Add("Content-Type", "application/json")
	jsonEncoder := json.NewEncoder(writer)
	err := jsonEncoder.Encode(response)
	DoPanicIfError(err)
}
