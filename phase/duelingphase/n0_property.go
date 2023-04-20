package duelingphase

import (
	"github.com/littletrainee/gunginotationgenerator/duelinground"
	"github.com/littletrainee/gunginotationgenerator/paircapture"
)

// 對弈期間的物件
type Duel struct {
	RoundList []duelinground.DuelRound
	Captured  []paircapture.PairCapture
	End       bool
}
