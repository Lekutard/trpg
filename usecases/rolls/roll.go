package rolls

import (
	"github.com/trpg/entities"
	"github.com/trpg/infrasteucture/db"
	"github.com/trpg/usecases/logging"
	"github.com/trpg/usecases/ports"
)

//RollUsecase Roll Usecase
type RollUsecase struct {
	DB      *db.DB
	Logging logging.Logging
}

//RollGeneral 一般行為を判定する
func (usecase *RollUsecase) RollGeneral(input *ports.GeneralRollInputPort) *ports.GeneralRollOutputPort {
	//6面体ダイスを二つ用意する
	dices := []entities.Dice{entities.Dice{Aspects: 6}, entities.Dice{Aspects: 6}}

	//目標値に補正を足す
	targetValue := input.TargetValue + input.Ajust

	//ダイスを振り、出目の合計値を出す
	var sum int
	for _, dice := range dices {
		sum = sum + dice.Throw()
	}

	//合計値が目標値以上なら成功。差を返す
	return &ports.GeneralRollOutputPort{
		Success:    (targetValue <= sum || sum == 12) && sum != 2,
		Critical:   sum == 12,
		Fumble:     sum == 2,
		Difference: sum - targetValue,
	}
}
