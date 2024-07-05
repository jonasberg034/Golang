package main

import "fmt"

func main() {

	menu := map[string]float64{

		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps

	for k, v := range menu {
		fmt.Println(k, "-", v)

	}
	// ints as key type

	phonebook := map[int]string{
		123456: "mario",
		78912:  "jonas",
	}

	fmt.Println(phonebook)
	fmt.Println(78912)

	phonebook[1233456] = "bowser"
	fmt.Println(phonebook)

	phonebook[45789] = "luigi"
	fmt.Println(phonebook)
}
