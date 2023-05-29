package main

import (
	"fmt"
	"net/http"
)

func signin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/signin" {
		r.ParseForm()
		tdn := r.FormValue("tdn")
		mk := r.FormValue("mk")

		sql := fmt.Sprintf("select * from taikhoan where tdn='%s' and mk='%s'", tdn, mk)
		db := connect("milkteashop")
		defer db.Close()
		rows, err := db.Query(sql)
		if err != nil {
			panic("loi sign in:" + err.Error())
		}
		i := 0
		var vaitro string
		var t string
		for ; rows.Next(); i++ {
			rows.Scan(&t, &t, &vaitro)
		}
		var kq string
		if i == 1 {
			kq = fmt.Sprintf("{'res':'%s'}", vaitro)
		} else {
			kq = "{'res':false}"
		}
		datajsonbyte := []byte(kq)

		w.Write(datajsonbyte)
	}

}
