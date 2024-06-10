package main

import (
	"fmt"
	"net/http"
)

// TODO:
//	* endpoint /note with GET
//  * endpoint /note with PUT

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("In /hello")
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error happenned: ", err.Error())
		return
	}
}
