package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type product struct {
	Masp  string `json:"masp"`
	Tensp string `json:"tensp"`
	Gia   int    `json:"gia"`
	Icon  string `json:"icon"`
}

func Rproduct(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/Rproduct" {
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
			sp.Masp = masp
			sp.Tensp = tensp
			sp.Gia = gia
			sp.Icon = iconencode

			fmt.Println("masp=" + masp + "::::::::::::")    // + tensp + string(gia) + iconencode
			fmt.Println("masp=" + sp.Masp + "::::::::::::") // + sp.tensp + string(sp.gia) + sp.icon

			data = append(data, sp)
		}
		datajsonbyte, err := json.Marshal(data)
		if err != nil {
			return
		}
		w.Write(datajsonbyte)
		fmt.Println(data)
		fmt.Println(string(datajsonbyte))
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
			data[i].Masp = masp
			data[i].Tensp = tensp
			data[i].Icon = iconencode
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

		fmt.Println(sp.Masp + ";;;;;;;" + sp.Tensp + ";;;;;;;;;;" + sp.Icon)

		if err != nil {
			panic(err)
		}

		//lay ma sp cu

		s := fmt.Sprintf("insert into sanpham(masp,tensp,gia,hinh) values('%s','%s','%d','%s')", sp.Masp, sp.Tensp, sp.Gia, sp.Icon)
		result := exe(s)
		var kq string
		if result {
			kq = "{'res':true}"
		} else {
			kq = "{'res':false}"
		}
		datajsonbyte, err := json.Marshal(kq)
		if err != nil {
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
		s := fmt.Sprintf("UPDATE sanpham SET tensp = '%s', gia = '%d', icon = '%s' WHERE masp = '%s' ", sp.Tensp, sp.Gia, sp.Icon, sp.Masp)
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
