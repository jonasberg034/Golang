package main

import "fmt"

func main() {

	// strings

	var nameOne string = "jonas"
	var nameTwo = "lille"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "diana"
	nameThree = "berg"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"

	fmt.Println(nameFour)

	// ints

	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits  & memory

	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint16 = 10  // sadece pozitif sayilar 0 to 255

	// floats

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 23.562
	scoreThree := 23.563

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}
