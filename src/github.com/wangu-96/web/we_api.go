package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: index")

}

func main() {

	http.HandleFunc("/", index)
	fmt.Println("Starting server at port 3000...")
	http.ListenAndServe(":3000", nil)

}
