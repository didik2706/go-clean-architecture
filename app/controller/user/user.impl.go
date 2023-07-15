package controller

import (
	usecase "latihan-restful-api-2/app/usecase/user"
	"latihan-restful-api-2/pkg"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) UserController {
	return &UserControllerImpl{
		UserUsecase: userUsecase,
	}
}

func (controller *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	request := usecase.UserUsecaseRequestCreate{}
	pkg.RequestParse(r, &request)

	data 		 := controller.UserUsecase.Create(r.Context(), request)
	response := pkg.ApiResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: data,
	} 

	pkg.SendResponse(w, response)
}

func (controller *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	request := usecase.UserUsecaseRequestUpdate{}
	pkg.RequestParse(r, &request)

	id, _ := strconv.Atoi(p.ByName("user_id"))
	request.Id = id

	data 		 := controller.UserUsecase.Update(r.Context(), request)
	response := pkg.ApiResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: data,
	}

	pkg.SendResponse(w, response)
}

func (controller *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("user_id"))

	controller.UserUsecase.Delete(r.Context(), id)
	response := pkg.ApiResponse{
		Code: http.StatusOK,
		Status: "OK",
	}

	pkg.SendResponse(w, response)
}

func (controller *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	data 		 := controller.UserUsecase.FindAll(r.Context())
	response := pkg.ApiResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: data,
	}

	pkg.SendResponse(w, response)
}

func (controller *UserControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("user_id"))

	data 		 := controller.UserUsecase.FindById(r.Context(), id)
	response := pkg.ApiResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: data,
	}

	pkg.SendResponse(w, response)
}

