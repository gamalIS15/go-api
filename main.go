package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"go-api/app"
	"go-api/controller"
	"go-api/helper"
	"go-api/middleware"
	"go-api/repository"
	"go-api/service"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.AuthMiddleware{router},
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
