package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type order struct {
	Madh        string `json:"madh"`
	Masp        string `json:"masp"`
	Thoihangiao string `json:"thoihangiao"`
	Diachi      string `json:"diachi"`
	Sodienthoai string `json:"sdt"`
}

func COrder(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/COrder" {
		var od order
		reqbpdy := json.NewDecoder(r.Body)
		err := reqbpdy.Decode(&od)

		if err != nil {
			panic(err)
		}

		s := fmt.Sprintf("insert into donhang(madh,masp,thoihangiao,diachi,sdt) values('%s','%s','%s','%s','%s')", od.Madh, od.Masp, od.Thoihangiao, od.Diachi, od.Sodienthoai)
		result := exe(s)
		var kq string
		if result {
			kq = "{'res':true}"
		} else {
			kq = "{'res':false}"
		}
		datajsonbyte := []byte(kq)
		w.Write(datajsonbyte)

	}
}

func DOrder(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/DOrder" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "error:", http.StatusBadRequest)
			return
		}
		s := fmt.Sprintf("delete from donhang where madh='%s'", r.FormValue("masp"))
		result := exe(s)
		var kq string
		if result {
			kq = "{'res':true}"
		} else {
			kq = "{'res':false}"
		}
		datajsonbyte := []byte(kq)
		w.Write(datajsonbyte)
	}
}
