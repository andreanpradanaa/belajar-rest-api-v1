package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, requestBody interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(requestBody)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, webResponse interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	PanicIfError(err)
}
