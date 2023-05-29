package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func autoIncrease(previousid string, increaseAmount int) string {
	previousValue, _ := strconv.Atoi(previousid[2:])
	increasedValue := previousValue + increaseAmount
	return fmt.Sprintf("%s%03d", previousid[0:2], increasedValue)
}

func main() {
	fileserver := http.FileServer(http.Dir("./../FE"))
	http.Handle("/", fileserver)

	http.HandleFunc("/Rproduct", Rproduct)
	http.HandleFunc("/Cproduct", Cproduct)
	http.HandleFunc("/Uproduct", Uproduct)
	http.HandleFunc("/Dproduct", Dproduct)

	http.HandleFunc("/COrder", COrder)
	http.HandleFunc("/DOrder", DOrder)

	http.HandleFunc("/searchproduct", searchproduct)
	http.HandleFunc("/signin", signin)
	fmt.Print("Starting server at port 1707\n")
	if err := http.ListenAndServe(":1707", nil); err != nil {
		log.Fatal(err)
	}

}
