package rolls

import (
	"testing"

	"github.com/trpg/usecases/ports"
)

func TestRollGeneral(t *testing.T) {
	usecase := RollUsecase{}

	// 通常の成功ケース
	t.Run("Rroll_general", func(t *testing.T) {

		input := &ports.GeneralRollInputPort{
			TargetValue: 6,
			Ajust:       6,
		}

		// テスト対象のビジネスロジック呼び出し
		usecase.RollGeneral(input)
	})
}
