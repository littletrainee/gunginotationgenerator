package duelingphase

import (
	"fmt"

	"github.com/littletrainee/ClearConsole"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/gunginotationgenerator/komahandler"
)

// 對弈期間的迴圈
func (d *DuelingPhase) Loop(komaHolder komahandler.KomaHandler,
	Level level.Level) {
	var round int = 1
	for !d.End {
		fmt.Printf("第 %d 巡\n\n", round)
		d.turn(komaHolder, round, Level)
		if d.End {
			break
		} else {
			round++
		}
		ClearConsole.ClearConsole()
	}
}
