package arrangementphase

import (
	"fmt"

	"github.com/littletrainee/ClearConsole"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/gunginotationgenerator/komahandler"
)

// 布置陣地期間的迴圈
func (a *ArrangementPhase) Loop(kh komahandler.KomaHandler, l level.Level) {
	var round int = 1
	for {
		fmt.Printf("第 %d 巡\n\n", round)
		a.turn(kh, round, l)
		ClearConsole.ClearConsole()
		if a.firstIsEnd && a.secondIsEnd {
			break
		} else {
			round++
		}
	}
}
