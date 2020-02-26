package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/main", func(rw http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("User")
		rw.Write([]byte(fmt.Sprintf("Welcome, %s", name)))
	})
	http.ListenAndServe(":8080", nil)
}
