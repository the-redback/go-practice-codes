package main

import "fmt"

func main(){
	var a [5]int

	a[4]=100;
	fmt.Println(a[4])
	fmt.Println(a)

	var x [5]float64

	x[0]=10
	x[1]=20
	x[2]=30
	x[3]=40
	x[4]=50

	total:=0.0

	for i:=0;i<len(x);i++{
		total+=x[i]
	}
	fmt.Println(total/float64(len(x)))


	y := [5]float64{ 98, 93, 77, 82, 83 }
	newTotal:=0.0

	for _, value:=range y{
		newTotal+=value
	}
	fmt.Println(newTotal/(float64(len(y))))

}
