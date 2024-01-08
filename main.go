package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":9000",
		Handler: nil,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
