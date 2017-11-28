package server

import (
	"net/http"
	"unit"
)

func Ng() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, res *http.Request) {
		resp.Write([]byte("hello world"))
	})

	//http.ListenAndServe(":8080", nil)
	unit.ListenAndServe(":8080", mux)
}
