package main

import "fmt"

func main() {
	var x map[string]int
	x = make(map[string]int)
	//x := make(map[string]int)   //declaration and Initialization

	x["key"] = 10
	x["k"] = 20
	x["b"] = 30
	x["a"] = 40

	fmt.Println(x)

	delete(x, "k")

	fmt.Println(x)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])
	fmt.Println(elements["F"])
	fmt.Println(elements["XX"])
	fmt.Println(elements)

	delete(elements, "H")

	fmt.Println(elements)

	name, ok := x["Un"]
	fmt.Println(name, ok)

	if name, ok := elements["O"]; ok {
		fmt.Println(name, ok)
	}

	elements2 := map[string]string{
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B": "Boron",
		"C": "Carbon",
		"N": "Nitrogen",
		"O": "Oxygen",
		"F": "Fluorine",
		"Ne": "Neon",
	}

	fmt.Println(elements2)

}