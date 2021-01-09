package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	post := Post{
		Id:      "1",
		Content: "Hello World!",
		Author: Author{
			Id:   "2",
			Name: "da ming",
		},
		Comments: []Comment{
			0: {
				Id: "3",
				Content: "Year year year year year!",
				Author: Author{
					Id:   "4",
					Name: "xiao ming",
				},
			},
			1: {
				Id: "5",
				Content: "How are you today!",
				Author: Author{
					Id:   "6",
					Name: "da hong",
				},
			},
		},
	}

	output, err := xml.MarshalIndent(&post, "", "\t")//Marshal结果相同，但不会格式化
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}

	err = ioutil.WriteFile("post.xml", []byte(xml.Header + string(output)), 0644)//前者会生成XML声明//err = ioutil.WriteFile("post.xml", output, 0644)
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
		return
	}
}
