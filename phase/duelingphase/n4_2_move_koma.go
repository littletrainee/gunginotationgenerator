package duelingphase

import (
	"fmt"
	"strconv"

	"github.com/littletrainee/gunginotationgenerator/capture"
	"github.com/littletrainee/gunginotationgenerator/enum/firstandsecondmove"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/gunginotationgenerator/enum/target"
	"github.com/littletrainee/gunginotationgenerator/komahandler"
	_phase "github.com/littletrainee/gunginotationgenerator/phase"
)

// 移居後確認是控制還是俘獲
func moveKoma(komaHolder komahandler.KomaHandler,
	firstorsecond firstandsecondmove.FirstAndSecondMove,
	level level.Level) capture.Capture {
	var (
		key     string
		capture capture.Capture = capture.Capture{}
		num     int64
		err     error
	)
	for {
		fmt.Print("是否俘獲對方的駒?(y/n))")
		fmt.Scanln(&key)
		if key != "y" && key != "n" {
			fmt.Println("輸入錯誤請重新輸入!")
			continue
		} else {
			break
		}
	}
	if key == "y" {
		capture.CapturedBy = firstorsecond

		for {
			fmt.Print("數量是幾個?")
			fmt.Scanln(&key)
			num, err = strconv.ParseInt(key, 10, 64)
			if err != nil || num < 1 || num > 3 {
				fmt.Println("輸入錯誤請重新輸入!")
				continue
			} else {
				break
			}
		}

		for i := 0; i < int(num); i++ {
			var (
				可能的選擇 string = _phase.Print(komaHolder, firstorsecond, level, target.OPPONENTBOARD)
				選擇的數字 int
				錯誤    error
			)
			for {
				fmt.Print("輸入選擇的駒：")
				fmt.Scanln(&key)
				選擇的數字, 錯誤 = strconv.Atoi(key)
				if 錯誤 != nil || 選擇的數字 > len(可能的選擇) || 選擇的數字 < 1 {
					fmt.Println("輸入錯誤請重新輸入!")
					continue
				} else {
					break
				}
			}
			key = _phase.ProbablyKoma(可能的選擇, 選擇的數字-1)
			fromBoardToOff(komaHolder, key, firstorsecond)
			capture.Koma += key
		}
	}
	return capture
}
