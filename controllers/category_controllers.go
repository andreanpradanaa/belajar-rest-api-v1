package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerInterfaces interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FIndAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
