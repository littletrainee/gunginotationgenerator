package duelingphase

import (
	"github.com/littletrainee/gunginotationgenerator/enum/firstandsecondmove"
	"github.com/littletrainee/gunginotationgenerator/komahandler"
)

// 移除對方的駒
func fromBoardToOff(komaHolder komahandler.KomaHandler, koma string,
	FirstAndSecond firstandsecondmove.FirstAndSecondMove) {
	if FirstAndSecond == firstandsecondmove.SECOND {
		komaHolder.FirstBoard[koma]--
	} else {
		komaHolder.SecondBoard[koma]--
	}
}
