package main

import "fmt"

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayGoodbye(n string) {
	fmt.Printf("Goodbye %v \n", n)

}

func cycleName(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func main() {
	sayGreeting("mario")
	sayGoodbye("luigi")
	cycleName([]string{"cloud", "tifa", "barret"}, sayGreeting)
}
