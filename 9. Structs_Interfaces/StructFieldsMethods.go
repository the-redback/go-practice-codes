package main

import (
	"math"
	"fmt"
)

type Circle struct{
	x,y,r float64
}

func (c *Circle) area() float64{   //Non pointers also works
	return math.Pi*c.r*c.r
}

func main() {
	c:=new(Circle)
	c.x=0
	c.y=0
	c.r=5

	fmt.Println(c)
	fmt.Println(c.area())


	cc := Circle{0, 0, 5}

	fmt.Println(cc)
	fmt.Println(cc.area())
}
