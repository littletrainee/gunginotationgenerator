package title

import (
	"image/color"

	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/gunginotationgenerator/komahandler"
)

// 回傳格式化後標題物件的屬性
func (t *Title) Format(kh komahandler.KomaHandler) string {
	var rt string = t.h3mark
	if kh.FirstColor == color.Black {
		rt += t.blackchess + "黑" + "</span>" + "棋方：" + t.black + "<br/>" +
			t.whitechess + "白" + "</span>" + "棋方：" + t.white + "<br/>"
	} else {
		rt += t.whitechess + "白" + "</span>" + "棋方：" + t.white + "<br/>" +
			t.blackchess + "黑" + "</span>" + "棋方：" + t.black + "<br/>"
	}
	rt += "對弈等級：" + level.GetLevelName(t.Level) + "\n\n"
	return rt
}
