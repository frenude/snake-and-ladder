package vo

import (
	"math/rand"
	"time"
)

type Dice int

func NewDice() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}
