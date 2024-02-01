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
    // Check for Cloudflare's CF-Connecting-IP header first
    cfIP := r.Header.Get("CF-Connecting-IP")
    if cfIP != "" {
        return cfIP
    }

    // Fallback to X-Forwarded-For, which might be set by Cloudflare or other proxies
    forwarded := r.Header.Get("X-FORWARDED-FOR")
    if forwarded != "" {
        ipList := strings.Split(forwarded, ",")
        return strings.TrimSpace(ipList[0]) 
    }

    // As a last resort, use the IP from the remote connection
    return strings.Split(r.RemoteAddr, ":")[0]
}
