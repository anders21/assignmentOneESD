package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	//http.Handle("/foo", fooHandler)

	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello test"+add(0, 9))
	})
/*
	log.Fatal(http.ListenAndServe(":8080", nil))
	*/


}
