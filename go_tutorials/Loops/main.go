package main

import (
	"fmt"
	"strings"
)

func main() {

	// x := 0
	// for x < 5 {
	//	fmt.Println("value of x is:", x)
	//	x++
	// for i := 0; i < 5; i++ {
	//	fmt.Println("value of i is:", i)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	greeting := "hello there friends"
	a := "hi"

	// for i := 0; i < len(names); i++ {
	//	fmt.Println("value of i is:", i)
	//}

	// for i := 0; i < len(names); i++ {
	//	fmt.Println(names[i])
	//	}

	// range: slice'ın her elemanı için bir döngü oluşturur.
	for _, value := range names { //index
		fmt.Printf("the value at index is %v \n", value)

	}

	for i := 0; i < len(strings.Split(greeting, " ")); i++ {

		fmt.Println("the value of i is:", a)
	}

}
