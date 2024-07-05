package main

import "fmt"

func main() {

	mybill := newBill("mario's bill")

	mybill.addItem("onion soup", 4.50)
	mybill.addItem("apple", 3.50)
	mybill.updateTip(10)

	fmt.Println(mybill.format())
}
