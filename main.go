package main

import (
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql" // import manual mysql
	"github.com/jalal-akbar/golang-restful-api/app"
	"github.com/jalal-akbar/golang-restful-api/controller"
	"github.com/jalal-akbar/golang-restful-api/helper"
	"github.com/jalal-akbar/golang-restful-api/middleware"
	"github.com/jalal-akbar/golang-restful-api/repository"
	"github.com/jalal-akbar/golang-restful-api/service"
)

func main() {

	validate := validator.New()
	// db := app.NewDB()
	// Controller
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryServiceImpl(categoryRepository, nil, validate)
	categoryController := controller.NewCategoryController(categoryService)
	// http Router impl
	router := app.NewRouter(categoryController)
	// Http Server Impl
	server := http.Server{
		Addr:    "localhost:3001",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
