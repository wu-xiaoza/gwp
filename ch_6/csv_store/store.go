package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id: 1, Content: "Hello World!", Author: "xiao ming"},
		Post{Id: 2, Content: "Bonjour Mundo!", Author: "da ming"},
		Post{Id: 3, Content: "Hola Mundo!", Author: "zhong ming"},
		Post{Id: 4, Content: "Greetings Earthlings!", Author: "xiao ming"},
	}

	//NewWriter returns a new Writer that writes to w.
	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		//Itoa is equivalent to FormatInt(int64(i), 10).
		//FormatInt returns the string representation of i in the given base, for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z' for digit values >= 10.
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		//Write writes a single CSV record to w along with any necessary quoting. A record is a slice of strings with each string being one field. Writes are buffered, so Flush must eventually be called to ensure that the record is written to the underlying io.Writer.
		//line写入writer
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	//Flush writes any buffered data to the underlying io.Writer. To check if an error occurred during the Flush, call Error.
	writer.Flush()

	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//NewReader returns a new Reader that reads from r.
	reader := csv.NewReader(file)
	//FieldsPerRecord is the number of expected fields per record. If FieldsPerRecord is positive, Read requires each record to have the given number of fields. If FieldsPerRecord is 0, Read sets it to the number of fields in the first record, so that future records must have the same field count. If FieldsPerRecord is negative, no check is made and records may have a variable number of fields.
	reader.FieldsPerRecord = -1
	//ReadAll reads all the remaining records from r. Each record is a slice of fields. A successful call returns err == nil, not err == io.EOF. Because ReadAll is defined to read until EOF, it does not treat end of file as an error to be reported.
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
