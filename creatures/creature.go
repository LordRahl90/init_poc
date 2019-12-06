package creatures

import (
	"fmt"
	"trial_test/randomness"
)

var creatures = []string{"ape", "lion", "panda", "leopard", "bear"}

func init() {
	fmt.Println("Creature method init is here")
}

//GiveRandom - Test the randomness
func GiveRandom() (creature string) {
	i := randomness.GiveRandNum(len(creatures))
	creature = creatures[i]
	return
}
