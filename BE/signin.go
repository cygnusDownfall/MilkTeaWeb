package main

import (
	"fmt"
	"log"
	"net/http"
)

func signin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/signin" {
		fmt.Println("chay")
	}
	log.Fatal("err")

}
