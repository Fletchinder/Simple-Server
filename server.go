package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ip := getIP(w, r)
	fmt.Fprintf(w, ip)
}

func getIP(w http.ResponseWriter, r *http.Request) string {
	IP := r.Header.Get("X-REAL-IP")
	if IP != "" {
		return IP
	} else {
		return "No IPV4"
	}
}
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
