package main

import "fmt"

func main() {
	var(
		a=4
		b='c'
		c="Hello"
	)
	fmt.Println(a,b,c)

	var input float64
	fmt.Println("Enter a Number: ")
	fmt.Scanf("%f",&input)

	output:=input*2
	fmt.Println(output)
}
