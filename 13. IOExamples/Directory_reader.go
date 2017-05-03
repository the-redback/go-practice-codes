package main

import (
	"os"
	"fmt"
)

func main() {
	dir, err := os.Open(".") //directory string

	if err!=nil{
		fmt.Println(err)
		return
	}

	defer dir.Close()

	fileInfos, err:=dir.Readdir(-1)

	for _,fi :=range fileInfos{
		fmt.Println(fi.Name())
	}
}
