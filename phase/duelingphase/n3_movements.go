package duelingphase

import (
	"fmt"
	"strconv"

	"github.com/littletrainee/gunginotationgenerator/capture"
	"github.com/littletrainee/gunginotationgenerator/enum/firstandsecondmove"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/gunginotationgenerator/enum/phase"
	"github.com/littletrainee/gunginotationgenerator/enum/target"
	"github.com/littletrainee/gunginotationgenerator/komaholder"
	_phase "github.com/littletrainee/gunginotationgenerator/phase"
	"github.com/littletrainee/gunginotationgenerator/position"
)

func (d *Duel) movements(komaHolder komaholder.KomaHolder, firstOrSecond firstandsecondmove.FirstAndSecondMove, 階級 level.Level) (string, capture.Capture) {
	var (
		位置      position.Position
		駒       string
		capture capture.Capture
		Is新     string
	)
ReEnterRCL:
	位置 = _phase.ColumnRowDan(firstOrSecond, 階級, phase.DUELING_PHASE)
	if 位置.Empty() {
		return "まげました", capture
	}

	for {
		var 確認列表 = _phase.Print(komaHolder, firstOrSecond, 階級, target.UNCAPTURED)
		if 階級 == level.ELEMENTARY || 階級 == level.BEGINNER {
			fmt.Println("注意：帥不可以堆疊在其他棋子上方！")
		}

		fmt.Print("請輸入選擇的駒或返回(b/B)重新選擇位置：")
		fmt.Scanln(&駒)
		if 駒 == "b" || 駒 == "B" {
			goto ReEnterRCL
		}
		選擇的數字, 錯誤 := strconv.Atoi(駒)
		if 錯誤 != nil {
			fmt.Println("輸入錯誤請重新輸入!")
			continue
		}
		駒 = _phase.GetString(確認列表, 選擇的數字-1)
		// 確認駒台是否有目標駒
		Is新 = arata(komaHolder, 駒, firstOrSecond)
		// 有目標駒的話執行新的動作，其他的都是移駒
		if Is新 == "新" {
			_phase.FromKomaDaiToBoard(komaHolder, 駒, firstOrSecond)
		} else {
			capture = moveKoma(komaHolder, firstOrSecond, 階級)
		}
		break

	}

	return 位置.Format() + 駒 + Is新, capture
}
