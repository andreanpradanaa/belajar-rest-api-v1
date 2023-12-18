package main

import (
	"belajar-rest-api/config"
	"belajar-rest-api/controllers"
	"belajar-rest-api/exception"
	"belajar-rest-api/middleware"
	"belajar-rest-api/repository"
	"belajar-rest-api/services"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"

	"github.com/julienschmidt/httprouter"
)

func main() {

	DB := config.NewDB()
	validate := validator.New()
	categoryRepo := repository.NewCategoryRepo()
	categoryService := services.NewCategoryService(categoryRepo, DB, validate)
	categoryController := controllers.NewCategoryController(categoryService)

	router := httprouter.New()
	router.GET("/api/v1/categories", categoryController.FIndAll)
	router.GET("/api/v1/categories/:categoryId", categoryController.FindById)
	router.POST("/api/v1/categories", categoryController.Create)
	router.PUT("/api/v1/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/v1/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("server running", server.ListenAndServe())
}
