package randomness

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Local().UnixNano())
	fmt.Println("Random Init is started here")
}

//GiveRandNum - func
func GiveRandNum(max int) (num int) {
	num = rand.Intn(max)
	return
}
