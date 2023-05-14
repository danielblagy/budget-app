package handler

import (
	"net/http"

	"github.com/danielblagy/budget-app/intenal/service/users"
)

type Handler interface {
	Greet(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	usersService users.Service
}

func NewHandler(usersService users.Service) Handler {
	return &handler{
		usersService: usersService,
	}
}
