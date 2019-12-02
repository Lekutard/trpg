package rolls

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/trpg/infrasteucture/db"
	"github.com/trpg/interfaces/api"
	"github.com/trpg/usecases/logging"
	"github.com/trpg/usecases/ports"
	"github.com/trpg/usecases/rolls"
)

//RollController Roll Controller
type RollController struct {
	Usecase *rolls.RollUsecase
}

//NewRollController Create New RollController
func NewRollController(db *db.DB, logging logging.Logging) *RollController {
	return &RollController{
		Usecase: &rolls.RollUsecase{
			DB:      db,
			Logging: logging,
		},
	}
}

// RollGeneral 一般判定の結果を返す
// @Summary 一般判定の結果を返す
// @Accept  json
// @Produce  json
// @Param target_value query int true "目標値"
// @Param ajust query int true "補正値。足すと難しくなり、引くと簡単になる"
// @Success 200 {object} ports.GeneralRollOutputPort
// @Header 200 {string} Token "qwerty"
// @Failure 500 {object} api.ErrorResponseObject
// @Router /rolls/general [get]
func (controller *RollController) RollGeneral(c *gin.Context) {
	targetValue, err := strconv.Atoi(api.GetQuery(c, "target_value"))
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponseObject{
			Message: err.Error(),
		})
		return
	}

	ajust, err := strconv.Atoi(api.GetQuery(c, "ajust"))
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponseObject{
			Message: err.Error(),
		})
		return
	}

	input := &ports.GeneralRollInputPort{
		TargetValue: targetValue,
		Ajust:       ajust,
	}

	output := controller.Usecase.RollGeneral(input)
	/*if uerr != nil {
		c.JSON(api.GetErrorResponse(uerr))
		return
	}*/

	c.JSON(http.StatusOK, output)
}
