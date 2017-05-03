package main

import "fmt"
import "./math"


func main(){
	xs:= []float64{1,2,3,4}
	avg := math.Average(xs)

	fmt.Println(avg)
}

//This project needs to be in go/src folder, because to install math.go, GOPATH is required.
// and GOPATH is  set to go/

//go run main.go