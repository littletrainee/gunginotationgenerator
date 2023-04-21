package arrangementphase

import (
	"github.com/littletrainee/gunginotationgenerator/arrangementround"
	"github.com/littletrainee/gunginotationgenerator/constant"
)

// 回傳格式化後的布陣期間物件的屬性
func (a *ArrangementPhase) Format() string {
	var rt string = "<center>\n\n"

	if len(a.roundList) > constant.MAXROW {
		var temp [][]arrangementround.ArrangementRound
		for i := 0; i < len(a.roundList); i += 10 {
			end := i + 10
			if end > len(a.roundList) {
				end = len((a.roundList))
			}
			temp = append(temp, a.roundList[i:end])
		}

		templength := len(temp)
		for i := 0; i < templength; i++ {
			rt += "|巡數|First|Second"

		}
		rt += "|\n"
		lastSliceIndex := 0
		for i := 0; i < 10; i++ {
			for j := 0; j < templength; j++ {
				if len(temp[j]) == len(temp[0]) {
					rt += temp[j][i].Format()
				} else if lastSliceIndex < len(temp[len(temp)-1]) {
					rt += temp[j][lastSliceIndex].Format()
					lastSliceIndex++
				}
			}
			rt += "|\n"
		}
	} else {
		rt += "|巡數|First|Second|\n"
		rt += "|---|---|---|\n"
		for i := 0; i < len(a.roundList); i++ {
			rt += a.roundList[i].Format() + "|\n"
		}
	}
	rt += "\n\n</center>\n"
	return rt
}
