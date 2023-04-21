package arrangementphase

import (
	"fmt"

	"github.com/littletrainee/gunginotationgenerator/arrangementround"
	"github.com/littletrainee/gunginotationgenerator/enum/firstandsecondmove"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/gunginotationgenerator/komahandler"
)

// 每一巡
func (a *ArrangementPhase) turn(kh komahandler.KomaHandler, index int, l level.Level) {
	var first string
	var second string
	if !a.firstIsEnd {
		first = a.movements(kh, firstandsecondmove.FIRST, l)
		if first == "済み" {
			a.firstIsEnd = true
		}
	} else {
		first = ""
	}
	fmt.Println()
	if !a.secondIsEnd {
		second = a.movements(kh, firstandsecondmove.SECOND, l)
		if second == "済み" {
			a.secondIsEnd = true
		}
	} else {
		first = ""
	}
	a.roundList = append(a.roundList, arrangementround.ArrangementRound{Order: index, First: first, Second: second})
}
