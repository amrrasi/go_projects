package management

import (
	"http-examples/handler"
	"net/http"
	"time"
)

func Run() {
	mux := http.NewServeMux()
	mux.Handle("/users/", &handler.UserHandler{})

	Server := &http.Server{
		Addr:         "8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Handler:      mux,
	}

	err := Server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
