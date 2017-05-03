package main

import (
	"os"
	"fmt"
)

func main() {

	file,err:=os.Create("FromCoding.txt")

	if err !=nil {
		fmt.Println("error in line 13:",err)
		return
	}

	defer file.Close()

	file.WriteString("This file is created from coding")
}
