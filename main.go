package main

import (
	"fmt"
	"golang-server/router"

	"net/http"
)

func main() {
	// Add Method GET
	fmt.Println("Go Server Start")
	http.HandleFunc("/get", router.GetAPI)
	http.HandleFunc("/post", post)
	http.ListenAndServe(":8080", nil)
}
