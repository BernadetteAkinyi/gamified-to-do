package main

import (
	"fmt"
	"gamified-to-do/api"
	"net/http"
)

func main() {

	http.HandleFunc("/", api.HomePage)
	http.HandleFunc("/tasks", api.AddTask)

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
