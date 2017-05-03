package main

import (
	"os"
	"fmt"
)

func main() {

	file,err := os.Open("test.txt")

	if err!=nil{
		//handle the error here
		fmt.Println("in line 14:", err)
		return
	}
	defer file.Close() //file will be closed after executing the main function

	//get the file size
	stat,err := file.Stat()



	//read the file

	bs:= make([]byte,stat.Size())
	_,err=file.Read(bs)


	str := string(bs)

	fmt.Println(str)
}
