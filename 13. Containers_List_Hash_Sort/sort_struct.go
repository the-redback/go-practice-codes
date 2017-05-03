package main

import (
	"sort"
	"fmt"
)

type Person struct{
	Name string
	Age int
}

type ByName []Person

// Len(), Less(), Swap() are 3 interfaces required for sort.Sort()
func (this ByName) Len() int{
	return len(this)
}

func (this ByName) Less(i,j int) bool{
	return this[i].Name < this[j].Name  //Sorts according to Name
}

func (this ByName) Swap(i,j int) {
	this [i],this[j]= this[j],this[i]
}

func main() {
	Kids:=[]Person{
		{"jill",9},
		{"Jack",10},
	}
	sort.Sort(ByName(Kids))
	fmt.Println(Kids)
}