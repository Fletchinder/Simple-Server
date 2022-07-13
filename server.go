package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ip := GetRealIP(r)
	fmt.Fprintf(w, ip)
}

func GetRealIP(r *http.Request) string {
	IP := r.Header.Get("X-Real-IP")
	if IP == "" {
		IP = r.Header.Get("X-Forwarder-For")
	}
	if IP == "" {
		IP = r.RemoteAddr
	}
	if IP == "" {
		IP = "No IPV4 address"
	}
	return IP
}
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
