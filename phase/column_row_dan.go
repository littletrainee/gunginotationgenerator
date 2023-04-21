// 期間包
package phase

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/littletrainee/gunginotationgenerator/boardposition"
	"github.com/littletrainee/gunginotationgenerator/enum/firstandsecondmove"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	_phase "github.com/littletrainee/gunginotationgenerator/enum/phase"
)

// 回傳縱橫斷詢問的結果
func ColumnRowDan(firstorsecond firstandsecondmove.FirstAndSecondMove,
	l level.Level, state _phase.Phase) boardposition.BoatdPosition {
	var (
		position boardposition.BoatdPosition = boardposition.BoatdPosition{}
		key      string
		maxlayer int
		err      error
	)
	if l != level.ADVANCED {
		maxlayer = 2
	} else {
		maxlayer = 3
	}
	for {
		fmt.Print("請輸入" + firstandsecondmove.FirstOrSecond(firstorsecond) + "的縱橫段(e/E)離開本階段：")
		fmt.Scanln(&key)
		if key == "e" || key == "E" {
			position.SetEmpty()
			return position
		}
		_, err = strconv.ParseInt(key, 10, 64)
		if err != nil || len(key) != 3 || strings.Contains(key, "0") {
			fmt.Println("輸入錯誤請重新輸入!")
			continue
		}
		column, err := strconv.Atoi(string(key[1]))
		if state == _phase.ARRANGEMENT_PHASE {
			if firstorsecond == firstandsecondmove.FIRST {
				if err != nil || column < 7 || column > 9 {
					fmt.Println("輸入錯誤請重新輸入!")
					continue
				}
			} else {
				if err != nil || column < 1 || column > 3 {
					fmt.Println("輸入錯誤請重新輸入!")
					continue
				}
			}
		} else {
			if err != nil || column < 1 || column > 9 {
				fmt.Println("輸入錯誤請重新輸入!")
			} else {
				break
			}

		}

		layer, err := strconv.Atoi(string(key[2]))
		if err != nil || layer < 1 || layer > maxlayer {
			fmt.Println("輸入錯誤請重新輸入!")
			continue
		}
		fmt.Println("輸入的縱橫段為：" + key + "\n正確請再按一次，不正確請隨意輸入")
		var check string
		fmt.Scanln(&check)
		if check != "" {
			continue
		} else {
			break
		}
	}
	position.Column = string(key[0])
	position.Row = string(key[1])
	position.Dan = string(key[2])

	return position
}
