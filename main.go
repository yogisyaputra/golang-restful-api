package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"yogisyaputra/golang-restful-api/app"
	"yogisyaputra/golang-restful-api/controller"
	"yogisyaputra/golang-restful-api/helper"
	"yogisyaputra/golang-restful-api/repository"
	"yogisyaputra/golang-restful-api/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRespository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRespository,db,validate)
	categoryController := controller.NewCategoryController(categoryService)

	route := httprouter.New()

	route.GET("/api/categories",categoryController.FindAll)
	route.GET("/api/categories/:categoryId",categoryController.FindById)
	route.POST("/api/categories",categoryController.Create)
	route.PUT("/api/categories/:categoryId",categoryController.Create)
	route.DELETE("/api/categories/:categoryId",categoryController.Delete)


	server := http.Server{
		Addr: "localhost:3000",
		Handler: route,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
