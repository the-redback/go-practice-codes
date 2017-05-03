package main

import "fmt"

func first() {
	fmt.Println("1st")
	panic("Panic in function First")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	defer second()
	first()
}

