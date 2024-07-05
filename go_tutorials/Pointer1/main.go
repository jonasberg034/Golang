package main

import "fmt"

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffe"] = 2.99
}

func main() {

	// group A types --> strings, ints, bools, floats, arrays, structs: These are non-pointer values
	name := "tifa"

	name = updateName(name)

	fmt.Println(name)

	// group B types --> slices, maps, functions: These are pointer values.

	menu := map[string]float64{

		"soup": 4.99,
		"pie":  7.99,
	}
	updateMenu(menu)
	fmt.Println(menu)
}
