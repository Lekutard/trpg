package usecase

import (
	"app/domain"
)

type Roll struct {
	dice  domain.Dice
	count int
}

func GraspDices(aspects int, count int) *Roll {
	dice := domain.ChooseDice(aspects)
	return &Roll{dice: *dice, count: count}
}

func (roll Roll) ThrowDices() (int, []int) {
	var sum int
	var detail []int
	for i := 0; i < roll.count; i++ {
		result := roll.dice.ThrowDice()
		detail = append(detail, result)
		sum = sum + result
	}
	return sum, detail
}
