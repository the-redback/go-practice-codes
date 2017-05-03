package main

import (
	"hash/crc32"
	"fmt"
)

func main() {

	h:=crc32.NewIEEE()
	h.Write([]byte("test"))
	v:=h.Sum32()
	fmt.Println("Hashed value:",v)

	h.Write([]byte("text"))
	v=h.Sum32()
	fmt.Println("Hashed value:",v)
}
