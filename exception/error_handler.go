package exception

import (
	"belajar-rest-api/helper"
	"belajar-rest-api/model/web"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	fmt.Println(err)
	if validationErrors(writer, request, err) {
		return
	}

	if notFoundError(writer, request, err) {
		fmt.Println("oooo")
		return
	}

	internalServerError(writer, request, err)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	}

	return false
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {

	fmt.Println("masuk sini")
	exception, ok := err.(NotFoundError)
	fmt.Println(ok, exception)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	}

	return false
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
