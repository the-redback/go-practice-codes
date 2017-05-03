package main

import "fmt"

type Person struct{
	name string
}

func (p *Person) Talk(){
	fmt.Println("The name is",p.name)
}

type Android struct{
	Person
	Model string
}

func main() {
	a:=new(Android)

	a.Person.name="Bond"
	a.Person.Talk()

	a.name="James Bond"
	a.Talk()

}
