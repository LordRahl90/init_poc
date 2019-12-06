package randomness

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("Random Init is started here")
}

//GiveRandNum - func
func GiveRandNum(max int) (num int) {
	num = rand.Intn(max)
	return
}
