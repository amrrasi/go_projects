package examples

import (
	"fmt"
	"net/http"
	"time"
)

type Testhandler struct {
}

func CreateServerWithMux() {

	server1 := http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Handler:      Testhandler{},
	}

	err1 := server1.ListenAndServe()
	if err1 != nil {
		panic(err1)
	}
}

func (h Testhandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
