package main

import (
	"io/ioutil"
	"hash/crc32"
	"fmt"
)

func getHash(filename string) (uint32,error){

	bs,err:=ioutil.ReadFile(filename)  //Corrected code of GoBook
	if err!=nil {
		return 0,err
	}

	h:=crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(),nil
}

func main() {

	h1,err :=getHash("test1.txt")

	if err != nil {
		fmt.Println("Error in test1",err)
		return
	}

	h2,err := getHash("test2.txt")

	if err!=nil {
		fmt.Println("Error in test2",err)
		return
	}

	fmt.Println(h1,h2,h1==h2)

}
