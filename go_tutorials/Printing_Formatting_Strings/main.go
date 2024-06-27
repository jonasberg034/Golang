package main

import "fmt"

func main() {

	age := 35
	name := "Jonas"

	// Print
	fmt.Print("hello, \n")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	// Println

	fmt.Println("hello ninjas!")
	fmt.Println("goodbye ninjas")
	fmt.Println("my name is:", name, "and my age is:", age)

	// Printf (Formatted strings) %_ = formatted specifier

	fmt.Printf("my name is %v and my age is: %v \n", name, age)
	fmt.Printf("my name is %q and my age is: %q \n", name, age)
	//fmt.Printf("age is type of %T", age)
	fmt.Printf("you scored is %f points \n", 225.55)
	fmt.Printf("you scored is %0.1f points \n", 225.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)
}
