package main

import (
	"latihan-restful-api-2/app/routes"
	"latihan-restful-api-2/config"
	"latihan-restful-api-2/exception"
	"latihan-restful-api-2/pkg"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// initialize DB
	db := config.NewDB()
	defer db.Close()

	// initialize router
	router := httprouter.New()

	// initialize validator
	validate := validator.New()

	// list routes
	routes.NewUserRoutes(router, db, validate)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	pkg.PanicIfError(err)
}