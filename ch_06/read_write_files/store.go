package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello World!\n")
	//WriteFile writes data to a file named by filename. If the file does not exist, WriteFile creates it with permissions perm (before umask); otherwise WriteFile truncates it before writing, without changing permissions.
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("data1")
	fmt.Print(string(read1))

	//Create creates or truncates the named file. If the file already exists, it is truncated. If the file does not exist, it is created with mode 0666 (before umask). If successful, methods on the returned File can be used for I/O; the associated file descriptor has mode O_RDWR. If there is an error, it will be of type *PathError.
	file1, _ := os.Create("data2")
	fmt.Println(file1)
	defer file1.Close()

	//Write writes len(b) bytes to the File. It returns the number of bytes written and an error, if any. Write returns a non-nil error when n != len(b).
	//这一步data2写入数据
	bytes, _ := file1.Write(data)
	fmt.Println(file1)
	fmt.Printf("Worte %d bytes to file\n", bytes)

	file2, _ := os.Open("data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))
}
