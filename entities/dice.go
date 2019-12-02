package entities

import (
	"math/rand"
	"time"
)

//Dice ダイスオブジェクト
type Dice struct {
	Aspects int //面体
}

//Throw ダイスが回転し、出目が変わる
func (dice Dice) Throw() int {
	time.Sleep(1) //1ナノ秒待つ
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(dice.Aspects-1) + 1
}
