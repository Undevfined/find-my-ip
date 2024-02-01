package main

import (
	"find-my-ip/handlers"
	"net/http"
)

func main() {
    http.HandleFunc("/", handlers.IPHandler)
    http.ListenAndServe(":80", nil)
}
