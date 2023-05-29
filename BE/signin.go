package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type taikhoan struct {
	Tentk string `json:"tentk"`
	Mk    string `json:"mk"`
}

func signin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/signin" {
		var tk taikhoan
		fmt.Println(r.Body)
		reqbpdy := json.NewDecoder(r.Body)
		err := reqbpdy.Decode(&tk)

		if err != nil {
			panic(err)
		}
		sql := fmt.Sprintf("select * from taikhoan where tentk='%s' and mk='%s'", tk.Tentk, tk.Mk)

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
			kq = "{\"res\":true}"
		} else {
			kq = "{\"res\":false}"
		}
		datajsonbyte := []byte(kq)

		w.Write(datajsonbyte)
	}

}
