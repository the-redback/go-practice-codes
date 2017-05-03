package main

import (
	"io/ioutil"
	"fmt"
)

func main() {

	bs,err := ioutil.ReadFile("test.txt")

	if err!= nil{
		fmt.Println("error is happended in line 13:",err)
		return
	}

	str:= string(bs)
	fmt.Println(str)
}
