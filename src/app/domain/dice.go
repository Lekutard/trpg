package domain

import (
	"math/rand"
	"time"
)

type Dice struct {
	aspects int
}

func ChooseDice(aspects int) *Dice {
	return &Dice{aspects: aspects}
}

func (dice Dice) ThrowDice() int {
	time.Sleep(1) //1ナノ秒待つ
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(dice.aspects-1) + 1
}
