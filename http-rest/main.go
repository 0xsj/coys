package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := 3000
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!\n")
	})

	fmt.Println("Server is running on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
