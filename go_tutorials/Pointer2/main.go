package main

import (
	"fmt"
)

func updateName(x *string) {
	*x = "wedge"

}

func main() {
	name := "tifa"

	// updateName(name)

	// fmt.Println("memory address of name is: ", &name)

	m := &name //* m, name değişkeninin bellek adresini tutan bir pointer olur.
	// fmt.Println(m)
	// fmt.Println("memory address of m is: ", m)
	// fmt.Println("value at memory address:", *m) // *m: m pointer'ının işaret ettiği bellek adresindeki değeri alır ve bu değeri yazdırır

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)

}
