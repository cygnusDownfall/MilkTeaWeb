package main

import (
	"fmt"
	"log"
	"net/http"
)

type product struct {
	masp       string
	tensp      string
	iconencode string
}

func Rproduct(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/product" {
		fmt.Fprint(w, "wellcome!!!")

		return
	}
	log.Fatal("err")
}

func searchproduct(w http.ResponseWriter, r *http.Request) {

}

func Cproduct(w http.ResponseWriter, r *http.Request) {
	var sp product
	if r.URL.Path == "/addproduct" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "error:", http.StatusBadRequest)
			return
		}
		sp.masp = r.FormValue("masp")
		sp.tensp = r.FormValue("tensp")
		sp.iconencode = r.FormValue("io")
		return
	}

}

func Uproduct(w http.ResponseWriter, r *http.Request) {

}

func Dproduct(w http.ResponseWriter, r *http.Request) {

}
