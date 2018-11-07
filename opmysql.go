package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//默认不输入ipport的信息的话是默认的127.0.0.1
	db, err := sql.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/idverify_db?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare("insert idverify_device set uuid =? , device_name=?, create_time=?")
	checkErr(err)

	res, err := stmt.Exec("7778888", "测试数据库", "2019-10-12")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	stmt, err = db.Prepare("update idverify_device set device_name=? where id =?")
	checkErr(err)

	res, err = stmt.Exec("name_update", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	rows, err := db.Query("SELECT id, uuid, device_name as devicename FROM idverify_device")
	checkErr(err)

	for rows.Next() {
		var id int
		var uuid string
		var devicename string
		err = rows.Scan(&id, &uuid, &devicename)
		checkErr(err)
		fmt.Printf("id = %d -- uuid = %s -- devicename = %s\n", id, uuid, devicename)
	}
	stmt, err = db.Prepare("delete from idverify_device where id = ?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect,err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
