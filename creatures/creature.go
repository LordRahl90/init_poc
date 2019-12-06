package creatures

import (
	"fmt"
	"trial_test/randomness"
)

var creatures = []string{"hello", "world"}

func init() {
	fmt.Println("Creature method init is here")
}

//GiveRandom - Test the randomness
func GiveRandom() (creature string) {
	i := randomness.GiveRandNum(len(creatures))
	fmt.Printf("I is: %d\n", i)
	creature = "hello"
	return
}
