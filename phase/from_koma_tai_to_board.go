package phase

import (
	"github.com/littletrainee/gunginotationgenerator/enum/firstandsecondmove"
	"github.com/littletrainee/gunginotationgenerator/komahandler"
)

// 從駒台到棋盤
func FromKomaTaiToBoard(komaHolder komahandler.KomaHandler, koma string,
	firstorsecond firstandsecondmove.FirstAndSecondMove) {
	if firstorsecond == firstandsecondmove.FIRST {
		komaHolder.FirstKomaTai[koma]--
		_, 存在 := komaHolder.FirstBoard[koma]
		if 存在 {
			komaHolder.FirstBoard[koma]++
		} else {
			komaHolder.FirstBoard[koma] = 1
		}
	} else {
		komaHolder.SecondKomaTai[koma]--
		_, 存在 := komaHolder.SecondBoard[koma]
		if 存在 {
			komaHolder.SecondBoard[koma]++
		} else {
			komaHolder.SecondBoard[koma] = 1
		}
	}

}
