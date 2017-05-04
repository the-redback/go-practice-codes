package main

import "fmt"

func main() {
	var x string
	x = "first "
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)

	var z string = "hello"
	var y string = "hello"
	fmt.Println(z == y)


	dogsName := "Max"   //Shortcut of Initializing  and assigning value
	fmt.Println("My dog's name is", dogsName)
}
