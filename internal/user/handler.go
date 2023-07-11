package user

import (
	"firstGOProject/internal/handlers"
	"firstGOProject/pkg/logging"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	usersUrl = "/users"
	userUrl  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersUrl, h.GetList)
	router.GET(userUrl, h.GetUserByUUID)
	router.POST(usersUrl, h.CreateUsers)
	router.PUT(userUrl, h.UpdateUser)
	router.PATCH(userUrl, h.PartiallyUpdateUser)
	router.DELETE(userUrl, h.DeleteUser)
}

func (h *handler) GetList(writer http.ResponseWriter, r *http.Request, params httprouter.Params) {
	writer.Write([]byte(fmt.Sprintf("users")))
}

func (h *handler) GetUserByUUID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Write([]byte(fmt.Sprintf("user")))

}

func (h *handler) CreateUsers(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Write([]byte(fmt.Sprintf("Create user")))

}

func (h *handler) UpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Write([]byte(fmt.Sprintf("Update user")))

}

func (h *handler) PartiallyUpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Write([]byte(fmt.Sprintf("PartiallyUpdate user")))

}

func (h *handler) DeleteUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Write([]byte(fmt.Sprintf("DeleteUser")))

}
