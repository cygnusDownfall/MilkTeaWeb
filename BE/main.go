package main

import (
	"fmt"
	"log"
	"net/http"
)



func main() {

	fileserver := http.FileServer(http.Dir("./FE"))
	http.Handle("/", fileserver)
	http.HandleFunc("/Rproduct", Rproduct)
	http.HandleFunc("/Cproduct", Cproduct)
	http.HandleFunc("/Uproduct", Uproduct)
	http.HandleFunc("/Dproduct", Dproduct)

	http.HandleFunc("/signin", signin)
	fmt.Print("Starting server at port 1707\n")
	if err := http.ListenAndServe(":1707", nil); err != nil {
		log.Fatal(err)
	}

}
