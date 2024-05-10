package main

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)


func TestConn(t *testing.T){
	db,err := sql.Open("sqlite3","test.db")
	if err != nil {
		fmt.Println("连接错误")
	}
	defer func ()  {_ = db.Close()}()
	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
	_,_ = db.Exec("CREATE TABLE User(Name text);")
	_,err =db.Exec("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam")
	if err != nil {
		fmt.Println("exec 错误")
		panic(err)
	}
	// fmt.Println(r1,"\n",r2,"\n",r3)
	row,err := db.Query("select Name from User LIMIT 2")
	var name string
	row.Next()
	var name2 string
	row.Scan(&name);
	row.Next()
	row.Scan(&name2);
	if err == nil{
		fmt.Println(name)
		fmt.Println(name2)
	}
	fmt.Println(err)
}