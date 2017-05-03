package main

import (
	"sort"
	"fmt"
)

func main() {
	xint:=[]int{2,1,6,5,0}
	xfloat:=[]float64{2.3,1.2,1.22,2.3,2.333}

	sort.Ints(xint)
	sort.Float64s(xfloat)

	fmt.Println(xint)
	fmt.Println(xfloat)

	xstring:=[]string{"aa","a","abc","cccc"}

	sort.Strings(xstring)

	fmt.Println(xstring)

}
