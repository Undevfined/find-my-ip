package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func IPHandler(w http.ResponseWriter, r *http.Request) {
    ip := getIP(r)
    fmt.Fprintf(w, "%s", ip)
}

func getIP(r *http.Request) string {
    forwarded := r.Header.Get("X-FORWARDED-FOR")
    if forwarded != "" {
        return forwarded
    }
    return strings.Split(r.RemoteAddr, ":")[0]
}
