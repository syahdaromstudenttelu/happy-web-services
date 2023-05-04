package helper

import "encoding/json"

func JoinResponse(responseOne any, responseTwo any) {
	resOneBytes, err := json.Marshal(responseOne)
	DoPanicIfError(err)

	err = json.Unmarshal(resOneBytes, &responseTwo)
	DoPanicIfError(err)
}
