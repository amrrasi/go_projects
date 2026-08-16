package examples

import (
	"net/http"
	"time"
)

func CreateServerWithMux() {

	mux := http.NewServeMux()

	mux.Handle("/google", http.RedirectHandler("https://www.google.com", 307))
	mux.Handle("/yahoo", http.RedirectHandler("https://www.yahoo.com", 307))
	server1 := http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Handler:      mux,
	}

	err1 := server1.ListenAndServe()
	if err1 != nil {
		panic(err1)
	}
}
