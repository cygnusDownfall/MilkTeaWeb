package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/home" {
		fmt.Fprint(w, "wellcome!!!")
		return
	}
	log.Fatal("err")

}

func signin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/signin" {
		signinpagehtml, err := os.ReadFile("FE/html/signin.html")
		if err != nil {
			log.Fatal("Khong tim thay trang dang nhap!")
		}
		fmt.Fprint(w, signinpagehtml)
		return
	}
	log.Fatal("err")

}

func main() {

	fileserver := http.FileServer(http.Dir("./FE"))
	http.Handle("/", fileserver)
	http.HandleFunc("/home", home)
	http.HandleFunc("/signin", signin)
	fmt.Print("Starting server at port 1707\n")
	if err := http.ListenAndServe(":1707", nil); err != nil {
		log.Fatal(err)
	}

}
