package helper

import "encoding/json"

func GetServiceDataResponse[D any](serviceResponse any, serviceResponseData D) {
	marshalledSvcUsersRes, err := json.Marshal(serviceResponse)
	DoPanicIfError(err)

	err = json.Unmarshal(marshalledSvcUsersRes, &serviceResponseData)
	DoPanicIfError(err)
}
