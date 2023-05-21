package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func connect(dbname string) *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/"+dbname)
	if err != nil {
		panic(err.Error())
	}

	return db
}

/*
func query(sqlQuery string) {

	db := connect("ql")
	defer db.Close()
	var (
		id   int
		name string
	)

	rows, err := db.Query(sqlQuery)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
}

func exe(sqlExe string) {

	db := connect("ql")
	defer db.Close()
	res, err := db.Exec(sqlExe)
	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ID: %d\n", lastId)
}
*/
