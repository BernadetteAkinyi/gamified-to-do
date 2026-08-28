package main

import (
	"fmt"
	"gamified-to-do/api"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks", api.HomePage)

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
