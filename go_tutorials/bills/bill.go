package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// * make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b

}

// * format the bills
func (b bill) format() string {
	fs := "Bill breakdown:  \n"
	var total float64 = 0

	//* list items

	for k, v := range b.items {
		fs += fmt.Sprintf("%v ...$%v \n", k+":", v)
		total += v
	}

	//* add tip

	fs += fmt.Sprintf("%v ...$%v \n", "tip:", b.tip)

	//* total

	fs += fmt.Sprintf("%v ...$%0.2f", "total:", total+b.tip)

	return fs

}

//* Update tip

func (b *bill) updateTip(tip float64) {
	b.tip = tip

}

//* add item

func (b *bill) addItem(names string, price float64) {
	b.items[names] = price

}

//* save bill

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("bill was saved")

}
