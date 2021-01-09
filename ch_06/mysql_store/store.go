package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	SNo       int
	SName     string
	SSex      string
	SBrithday string
	SClass    int
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:1990686992@/dbtest")
	if err != nil {
		panic(err)
	}
}

func main() {
	var students []Student
	rows, err := Db.Query("select * from student")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		student := Student{}
		err = rows.Scan(&student.SNo, &student.SName, &student.SSex, &student.SBrithday, &student.SClass)
		if err != nil {
			panic(err)
		}
		students = append(students, student)
	}
	for _, student := range students {
		fmt.Println(student)
	}
	rows.Close()
}