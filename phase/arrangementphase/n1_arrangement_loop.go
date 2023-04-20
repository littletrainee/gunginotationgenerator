package arrangementphase

import (
	"fmt"

	"github.com/littletrainee/ClearConsole"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/gunginotationgenerator/komaholder"
)

// 布置陣地期間的迴圈
func (s *SetUp) SetUpLoop(kh komaholder.KomaHolder, l level.Level) {
	var round int = 1
	for {
		fmt.Printf("第 %d 巡\n\n", round)
		s.turn(kh, round, l)
		ClearConsole.ClearConsole()
		if s.firstIsEnd && s.secondIsEnd {
			break
		} else {
			round++
		}
	}
}
