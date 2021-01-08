package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp password=1990686992 port=5433 dbname=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	//Scan copies the columns from the matched row into the values pointed at by dest. See the documentation on Rows.Scan for details. If more than one row matches the query, Scan uses the first row and discards the rest. If no row matches the query, Scan returns ErrNoRows.
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

//Create a new post.
func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	//Prepare creates a prepared statement for later queries or executions. Multiple queries or executions may be run concurrently from the returned statement. The caller must call the statement's Close method when the statement is no longer needed.
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	//QueryRow executes a prepared query statement with the given arguments. If an error occurs during the execution of the statement, that error will be returned by a call to Scan on the returned *Row, which is always non-nil. If the query selects no rows, the *Row's Scan will return ErrNoRows. Otherwise, the *Row's Scan scans the first selected row and discards the rest.
	//Scan copies the columns from the matched row into the values pointed at by dest. See the documentation on Rows.Scan for details. If more than one row matches the query, Scan uses the first row and discards the rest. If no row matches the query, Scan returns ErrNoRows.
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) Update() (err error) {
	//Exec executes a query without returning any rows. The args are for any placeholder parameters in the query.
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delate() (err error) {
	_, err = Db.Exec("delate from posts where id = $1", post.Id)
	return
}

func main() {
	post1 := Post{Content: "Hello World!", Author: "xiao ming"}
	post2 := Post{Content: "Hello World!", Author: "da ming"}
	post3 := Post{Content: "Hello World!", Author: "zhong ming"}
	post4 := Post{Content: "Hello World!", Author: "gao ming"}

	fmt.Println(post1)
	post1.Create()
	post2.Create()
	post3.Create()
	post4.Create()
	fmt.Println(post1)

	readPost, _ := GetPost(post1.Id)
	fmt.Println(readPost)

	readPost.Content = "gudebai"
	readPost.Author = "da hong"
	readPost.Update()

	posts, _ := Posts(4)
	fmt.Println(posts)

	readPost.Delate()
}

//
