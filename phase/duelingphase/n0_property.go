// 對弈期間包
package duelingphase

import (
	"github.com/littletrainee/gunginotationgenerator/duelinground"
	"github.com/littletrainee/gunginotationgenerator/paircapture"
)

// 對弈期間的物件
type DuelingPhase struct {
	RoundList []duelinground.DuelingRound
	Captured  []paircapture.PairCapture
	End       bool
}
