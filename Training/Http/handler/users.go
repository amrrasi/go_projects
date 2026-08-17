package handler

import (
	"fmt"
	"http-examples/model"
	"net/http"
)

type UserHandler struct {
}

func (h UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch {

	case r.Method == "GET" && len(r.URL.Query().Get("id")) > 0:
		GetUser(w, r)
		return
	case r.Method == "Get" && len(r.URL.Query().Get("id")) == 0:
		GetUserList(w, r)
		return
	case r.Method == "POST":
		CreateUser(w, r)
		return
	}

}

func CreateUser(u model.User, w http.ResponseWriter, r *http.Request) {
	authorizationKey := r.Header.Get("x-Authorization")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println(w, "Get user with id:", id)
}

func GetUserList(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Get userList")
}
