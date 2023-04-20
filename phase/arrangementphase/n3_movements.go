package arrangementphase

import (
	"fmt"
	"strconv"

	"github.com/littletrainee/gunginotationgenerator/enum/firstandsecondmove"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/gunginotationgenerator/enum/phase"
	"github.com/littletrainee/gunginotationgenerator/enum/target"
	"github.com/littletrainee/gunginotationgenerator/komaholder"
	_phase "github.com/littletrainee/gunginotationgenerator/phase"
	"github.com/littletrainee/gunginotationgenerator/position"
)

func (s *SetUp) movements(kh komaholder.KomaHolder,
	firstorsecond firstandsecondmove.FirstAndSecondMove, 階級 level.Level) string {
	var (
		位置 position.Position = position.Position{}
		駒  string
	)
ReEnterRCL:
	位置 = _phase.ColumnRowDan(firstorsecond, 階級, phase.ARRANGEMENT_PHASE)
	if 位置.Empty() {
		return "済み"
	}
	for {
		var 可能的選擇 string = _phase.Print(kh, firstorsecond, 階級, target.KOMATAI)
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
		駒 = _phase.GetString(可能的選擇, 選擇的數字-1)
		_phase.FromKomaDaiToBoard(kh, 駒, firstorsecond)
		break
	}

	return 位置.Format() + 駒
}
