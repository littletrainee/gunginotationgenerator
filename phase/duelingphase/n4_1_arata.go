package duelingphase

import (
	"fmt"

	"github.com/littletrainee/gunginotationgenerator/enum/firstandsecondmove"
	"github.com/littletrainee/gunginotationgenerator/komahandler"
)

// 確認是否從備排區移動到棋盤
func arata(komaHolder komahandler.KomaHandler, koma string,
	firstorsecond firstandsecondmove.FirstAndSecondMove) string {
	var key string
	if firstorsecond == firstandsecondmove.FIRST {
		if komaHolder.FirstKomaTai[koma] > 0 {
			for {
				fmt.Print(`是否要為"新加入"?(y/n)`)
				fmt.Scanln(&key)
				if key != "y" && key != "n" {
					fmt.Println("輸入錯誤請重新輸入!")
					continue
				} else {
					break
				}
			}
			if key == "y" {
				return "新"
			}
		}
	} else {
		if komaHolder.SecondKomaTai[koma] > 0 {
			for {
				fmt.Print(`是否要為"新加入"?(y/n)`)
				fmt.Scanln(&key)
				if key != "y" && key != "n" {
					fmt.Println("輸入錯誤請重新輸入!")
					continue
				} else {
					break
				}
			}
			if key == "y" {
				return "新"
			}
		}
	}
	return ""
}
