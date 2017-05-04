package main

import "fmt"

func main() {
	i:=1
	for i<=10{
		fmt.Println(i)
		i++;
	}

	for j:=1;j<=10;j++ {
		fmt.Println(j)
	}
	fmt.Println(i)
}
