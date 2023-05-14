package handler

import (
	"fmt"
	"net/http"
)

func (h handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.usersService.GetUsers(r.Context())
	if err != nil {
		fmt.Fprintf(w, "err: %s", err.Error())
		return
	}

	for _, user := range users {
		fmt.Fprintf(w, "%s --- %s --- %s\n", user.Username, user.Email, user.FullName)
	}
}
