package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
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
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML file:", err)
		return
	}

	var post Post
	xml.Unmarshal(xmlData, &post)

	// 	{ post}
	fmt.Println(post.XMLName)
	fmt.Println()
	// post
	fmt.Println(post.XMLName.Local)
	fmt.Println()
	//
	fmt.Println(post.XMLName.Space)
	fmt.Println()
	// 1
	fmt.Println(post.Id)
	fmt.Println()
	// Hello World!
	fmt.Println(post.Content)
	fmt.Println()
	// {2 Sau Sheong}
	fmt.Println(post.Author)
	fmt.Println()
	//     <content>Hello World!</content>
	//     <author id="2">Sau Sheong</author>
	//     <comments>
	//         <comment id="1">
	//             <content>Have a great day!</content>
	//             <author id="3">Adam</author>
	//         </comment>
	//         <comment id="2">
	//             <content>How are you today?</content>
	//             <author id="4">Betty</author>
	//         </comment>
	//     </comments>
	fmt.Println(post.Xml)
	fmt.Println()
	// [{1 Have a great day! {3 Adam}} {2 How are you today? {4 Betty}}]
	fmt.Println(post.Comments)
	fmt.Println()
	// 2
	fmt.Println(post.Author.Id)
	fmt.Println()
	// Sau Sheong
	fmt.Println(post.Author.Name)
	fmt.Println()
	// {1 Have a great day! {3 Adam}}
	fmt.Println(post.Comments[0])
	fmt.Println()
	// 1
	fmt.Println(post.Comments[0].Id)
	fmt.Println()
	// Have a great day!
	fmt.Println(post.Comments[0].Content)
	fmt.Println()
	//{3 Adam}
	fmt.Println(post.Comments[0].Author)
	fmt.Println()
	// {2 How are you today? {4 Betty}}
	fmt.Println(post.Comments[1])
	fmt.Println()
	// 2
	fmt.Println(post.Comments[1].Id)
	fmt.Println()
	// How are you today?
	fmt.Println(post.Comments[1].Content)
	fmt.Println()
	// {4 Betty}
	fmt.Println(post.Comments[1].Author)
}
