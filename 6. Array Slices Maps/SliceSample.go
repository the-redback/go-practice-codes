package main

import "fmt"

func main(){
	//var x []float64   //Slice is without index number
	x:=make([]float64,5)
	x[0]=100
	fmt.Println(x)


	arr := []float64{1,2,3,4,5}
	y := arr[0:3]   //Slice arr[low:high] 1,2,3

	fmt.Println(y)

	slice1:=[]int{11,22,33,44,55}
	slice2:=append(slice1,66,77)


	fmt.Println(slice1,slice2)

	slice3:=[]int{111,222,333,444}
	slice4:=make([]int,2)

	copy(slice4,slice3)

	fmt.Println(slice3,slice4)

	fmt.Println(len(slice1),len(slice2),len(slice3),len(slice4))




}
