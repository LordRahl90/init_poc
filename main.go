package main

import (
	"fmt"
	"trial_test/creatures"
)

func init() {
	fmt.Println("Main method init comes here")
}

func main() {
	c := creatures.GiveRandom()
	fmt.Println(c)
}
