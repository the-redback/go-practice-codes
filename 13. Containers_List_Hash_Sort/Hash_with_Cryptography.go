package main

import (
	"crypto/sha1"
	"fmt"
)

//sha1 and crc32 implements the hash.hash interface

//sha1 computes a 160 bit hash. but as there is no such data type to hold this number, a slice of 20 bytes is used
func main() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs:=h.Sum([]byte{})

	fmt.Println(bs)
}
