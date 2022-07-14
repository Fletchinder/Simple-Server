package main

import (
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ip := GetRealIP(r)
	WritetoFile(ip + "\n")
	http.Redirect(w, r, "https://www.youtube.com/watch?v=ZVQDHFgfssM", 301)
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
func WritetoFile(IP string) {
	f, err := os.OpenFile("ListofIPs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(IP + "\n"); err != nil {
		log.Println(err)
	}
}
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
