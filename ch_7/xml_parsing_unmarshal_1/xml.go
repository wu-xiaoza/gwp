package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	Xml     string   `xml:",innerxml"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}

	var post Post
	xml.Unmarshal(xmlData, &post)
	/* 	{{ post} 1 Hello World! {2 da ming}
	<content>Hello World!</content>
	<author id="2">da ming</author>
	} */
	fmt.Println(post)
	fmt.Println()

	fmt.Println(post.XMLName)     //{ post}
	fmt.Println()
	fmt.Println(post.Id)          //1
	fmt.Println()
	fmt.Println(post.Content)     //Hello World!
	fmt.Println()
	fmt.Println(post.Author)      //{2 da ming}
	fmt.Println()
	// <content>Hello World!</content>
	// <author id="2">da ming</author>
	fmt.Println(post.Xml)
	fmt.Println()
	fmt.Println(post.Author.Id)   //2
	fmt.Println()
	fmt.Println(post.Author.Name) //da ming
}
