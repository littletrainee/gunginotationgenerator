package duelingphase

import (
	"github.com/littletrainee/gunginotationgenerator/enum/firstandsecondmove"
	"github.com/littletrainee/gunginotationgenerator/komaholder"
)

// 移除對方的駒
func fromBoardToOff(komaHolder komaholder.KomaHolder, koma string,
	FirstAndSecond firstandsecondmove.FirstAndSecondMove) {
	if FirstAndSecond == firstandsecondmove.SECOND {
		komaHolder.FirstBoard[koma]--
	} else {
		komaHolder.SecondBoard[koma]--
	}
}
