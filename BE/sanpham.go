package main

import (
	"encoding/json"
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
		err := r.ParseForm()
		if err == nil {
			return
		}
		if r.FormValue("product") == "all" {
			var data []product
			db := connect("milkteashop")
			defer db.Close()
			row, err := db.Query("select * from sanpham")
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
				data[i].iconencode = iconencode
			}
			datajsonbyte, err := json.Marshal(data)
			if err == nil {
				return
			}
			w.Write(datajsonbyte)
		}
		return
	}
	log.Fatal("err")
}

func searchproduct(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/product" {
		err := r.ParseForm()
		if err == nil {
			return
		}
		if r.FormValue("product") == "all" {
			var data []product
			db := connect("milkteashop")
			defer db.Close()
			row, err := db.Query("select * from sanpham")
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
				data[i].iconencode = iconencode
			}
			datajsonbyte, err := json.Marshal(data)
			if err == nil {
				return
			}
			w.Write(datajsonbyte)
		}
		return
	}
	log.Fatal("err")
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
		sp.iconencode = r.FormValue("icon")

		result := exe("insert into sanpham(masp,tensp,icon) values('%s','%s','%s')")
		if result {
			w.Header().Add("result", "thanh cong")
		} else {
			w.Header().Add("result", "khong thanh cong")
		}

		return
	}

}

func Uproduct(w http.ResponseWriter, r *http.Request) {
	var sp product
	if r.URL.Path == "/addproduct" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "error:", http.StatusBadRequest)
			return
		}
		sp.masp = r.FormValue("masp")
		sp.tensp = r.FormValue("tensp")
		sp.iconencode = r.FormValue("icon")

		result := exe("insert into sanpham(masp,tensp,icon) values('%s','%s','%s')")
		if result {
			w.Header().Add("result", "thanh cong")
		} else {
			w.Header().Add("result", "khong thanh cong")
		}

		return
	}
}

func Dproduct(w http.ResponseWriter, r *http.Request) {
	var sp product
	if r.URL.Path == "/addproduct" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "error:", http.StatusBadRequest)
			return
		}
		sp.masp = r.FormValue("masp")
		sp.tensp = r.FormValue("tensp")
		sp.iconencode = r.FormValue("icon")

		result := exe("insert into sanpham(masp,tensp,icon) values('%s','%s','%s')")
		if result {
			w.Header().Add("result", "thanh cong")
		} else {
			w.Header().Add("result", "khong thanh cong")
		}

		return
	}
}
