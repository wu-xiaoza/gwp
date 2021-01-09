package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Nmae string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author`
}

func main() {
	post := Post{
		Id:      1,
		Content: "Hellow World!",
		Author: Author{
			Id:   2,
			Nmae: "Wang Fugui",
		},
		Comments: []Comment{
			{
				Id:      3,
				Content: "Are you ok!",
				Author:  "Lei Jun",
			}, {
				Id:      4,
				Content: "What the fuck!",
				Author:  "Wuai Guoren",
			},
		},
	}

	output, err := json.MarshalIndent(&post, "", "\t\t")// output, err := json.Marshal(&post)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	err = ioutil.WriteFile("post.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}
