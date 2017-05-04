package main

import "fmt"

func main(){
	fmt.Println("Hello there,\n\tThis is Maruf. \n\nThanx,\nMaruf" )
	fmt.Println("Hello there,\n\tThis is Maruf. \n\nThanx,\nMaruf"[1] )
	fmt.Println(len("Hello there,\n\tThis is Maruf. \n\nThanx,\nMaruf" ))
	fmt.Println(`Hello there,\n\tThis is Maruf. \n\nThanx,\nMaruf` )
	fmt.Println(`Hello there,\n\tThis is Maruf. \n\nThanx,\nMaruf`[1] )
	fmt.Println(len(`Hello there,\n\tThis is Maruf. \n\nThanx,\nMaruf`) )
	fmt.Println("Hello "+"World")
}
