package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type product struct {
	masp  string
	tensp string
	gia   int
	icon  string
}

func Rproduct(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/Rproduct" {
		fmt.Print("R")
		var data []product
		var sp product
		db := connect("milkteashop")
		defer db.Close()
		row, err := db.Query("select * from sanpham")
		if err != nil {
			panic(err.Error())
		}
		var masp, tensp, iconencode string
		var gia int
		for i := 0; row.Next(); i++ {
			err := row.Scan(&masp, &tensp, &gia, &iconencode)
			if err != nil {
				panic(err.Error())
			}
			sp.masp = masp
			sp.tensp = tensp
			sp.gia = gia
			sp.icon = iconencode

			fmt.Print(masp + tensp + string(gia) + iconencode)
			fmt.Print(sp.masp + sp.tensp + string(sp.gia) + sp.icon)

			data = append(data, sp)
		}
		datajsonbyte, err := json.Marshal(data)
		if err != nil {
			return
		}
		w.Write(datajsonbyte)
		fmt.Print(data)
		return
	}
}

func searchproduct(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/searchproduct" {

		var data []product
		db := connect("milkteashop")
		defer db.Close()
		row, err := db.Query("select * from sanpham where tensp=\"" + r.FormValue("tensp") + "\" ")
		if err != nil {
			panic(err.Error())
		}
		var masp, tensp, iconencode string
		for i := 0; row.Next(); i++ {
			err := row.Scan(&masp, &tensp, &iconencode)
			if err != nil {
				panic(err.Error())
			}
			data[i].masp = masp
			data[i].tensp = tensp
			data[i].icon = iconencode
		}
		datajsonbyte, err := json.Marshal(data)
		if err == nil {
			return
		}
		w.Write(datajsonbyte)

		return
	}
}

func Cproduct(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/Cproduct" {
		var sp product
		reqbpdy := json.NewDecoder(r.Body)
		err := reqbpdy.Decode(&sp)

		fmt.Println(sp.masp + sp.tensp + sp.icon)

		if err != nil {
			panic(err)
		}

		s := fmt.Sprintf("insert into sanpham(masp,tensp,gia,hinh) values('%s','%s','%d','%s')", sp.masp, sp.tensp, sp.gia, sp.icon)
		result := exe(s)
		var kq string
		if result {
			kq = "{'res':true}"
		} else {
			kq = "{'res':false}"
		}
		datajsonbyte, err := json.Marshal(kq)
		if err == nil {
			return
		}
		w.Write(datajsonbyte)
		return
	}

}

func Uproduct(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/Uproduct" {
		var sp product
		reqbpdy := json.NewDecoder(r.Body)
		err := reqbpdy.Decode(&sp)
		if err != nil {
			panic(err)
		}
		s := fmt.Sprintf("UPDATE sanpham SET tensp = '%s', gia = '%d', icon = '%s' WHERE masp = '%s' ", sp.tensp, sp.gia, sp.icon, sp.masp)
		result := exe(s)
		var kq string
		if result {
			kq = "{\"res\":true}"
		} else {
			kq = "{\"res\":false}"
		}
		datajsonbyte, err := json.Marshal(kq)
		if err == nil {
			return
		}
		w.Write(datajsonbyte)
		return
	}
}

func Dproduct(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/Dproduct" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "error:", http.StatusBadRequest)
			return
		}
		s := fmt.Sprintf("delete from sanpham where masp='%s'", r.FormValue("masp"))
		result := exe(s)
		var kq string
		if result {
			kq = "{'res':true}"
		} else {
			kq = "{'res':false}"
		}
		datajsonbyte, err := json.Marshal(kq)
		if err == nil {
			return
		}
		w.Write(datajsonbyte)
		return
	}
}
