package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("--- request received")

		for name, values := range r.Header {
			for _, value := range values {
				fmt.Println(name, value)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	http.ListenAndServe(":3000", server)
}
