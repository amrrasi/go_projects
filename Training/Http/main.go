package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	go func() {
		Server1 := &http.Server{
			Addr:         ":8080",
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
		}

		err := Server1.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	log.Fatal(http.ListenAndServe(":8090", nil))
}
