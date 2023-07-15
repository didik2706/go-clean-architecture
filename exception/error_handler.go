package exception

import (
	"latihan-restful-api-2/pkg"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) { return }
	if badRequestError(w, r, err) { return }

	internalServerError(w, r, err)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.WriteHeader(http.StatusNotFound)

		response := pkg.ApiResponse{
			Code: http.StatusNotFound,
			Status: "NOT FOUND",
			Data: exception.Error,
		}

		pkg.SendResponse(w, response)
		return true
	} else {
		return false
	}
}

func badRequestError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.WriteHeader(http.StatusBadRequest)

		response := pkg.ApiResponse{
			Code: http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: exception,
		}

		pkg.SendResponse(w, response)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.WriteHeader(http.StatusInternalServerError)

	response := pkg.ApiResponse{
		Code: http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data: err,
	}

	pkg.SendResponse(w, response)
}