package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			fmt.Println("Get POST request")
			w.Write([]byte("hello world dari post"))
			return
		}
		fmt.Println("Get non-POST request")
		w.Write([]byte("hello world"))
	})
	fmt.Println("Server running on localhost 8080")
	http.ListenAndServe(":8080", nil)
}