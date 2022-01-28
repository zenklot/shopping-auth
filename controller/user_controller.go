package controller

import (
	"encoding/json"
	"net/http"
	"shopping/model/web"
	"shopping/service"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Sigin(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}

}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateReq := web.UserCreateRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userCreateReq)
	if err != nil {
		panic(err)
	}

	userResponse := controller.UserService.Create(request.Context(), userCreateReq)

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(userResponse)
	if err != nil {
		panic(err)
	}
}

func (controller *UserControllerImpl) Sigin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userSignReq := web.UserSignRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userSignReq)
	if err != nil {
		panic(err)
	}

	userResponse := controller.UserService.Sign(request.Context(), userSignReq)

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(userResponse)
	if err != nil {
		panic(err)
	}
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}
