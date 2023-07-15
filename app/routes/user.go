package routes

import (
	"database/sql"
	controller "latihan-restful-api-2/app/controller/user"
	repository "latihan-restful-api-2/app/repository/user"
	usecase "latihan-restful-api-2/app/usecase/user"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func NewUserRoutes(route *httprouter.Router, db *sql.DB, validate *validator.Validate) *httprouter.Router {
	userRepository := repository.NewUserRepository()
	userUsecase		 := usecase.NewUserUsecase(userRepository, db, validate)
	userController := controller.NewUserController(userUsecase)
	
	route.GET("/api/users", userController.FindAll)
	route.GET("/api/users/:user_id", userController.FindById)
	route.POST("/api/users", userController.Create)
	route.PUT("/api/users/:user_id", userController.Update)
	route.DELETE("/api/users/:user_id", userController.Delete)

	return route
}