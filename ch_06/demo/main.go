package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte("qqqqq qqqqq!\n")
	fmt.Print(data)
	file1, _ := os.Create("data2")
	fmt.Println(file1)
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Println(file1)
	fmt.Printf("Worte %d bytes to file\n", bytes)
}
