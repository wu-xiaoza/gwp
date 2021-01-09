package main

import (
	"encoding/xml"
	"fmt"
	"os"
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
				Id:      "3",
				Content: "Year year year year year!",
				Author: Author{
					Id:   "4",
					Name: "xiao ming",
				},
			},
			1: {
				Id:      "5",
				Content: "How are you today!",
				Author: Author{
					Id:   "6",
					Name: "da hong",
				},
			},
		},
	}

	xmlFile, err := os.Create("post.xml")
	if err != nil {
		fmt.Println("Error creating XML file:", err)
		return
	}
	_, err = xmlFile.Write([]byte(xml.Header))//XML声明
	if err != nil {
		fmt.Println("Error writing XML.Header to file:", err)
	}
	
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding XML to file:", err)
		return
	}
}
