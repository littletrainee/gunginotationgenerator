package duelingphase

import (
	"github.com/littletrainee/gunginotationgenerator/constant"
	"github.com/littletrainee/gunginotationgenerator/duelinground"
)

// 回傳格式化後的對弈歧見物件的屬性
func (d *DuelingPhase) Format() string {
	var rt string = "<center>\n\n"
	if len(d.RoundList) > constant.MAXROW {
		var temp [][]duelinground.DuelingRound
		//分割為每欄10行
		for i := 0; i < len(d.RoundList); i += 10 {
			end := i + 10
			if end > len(d.RoundList) {
				end = len(d.RoundList)
			}
			temp = append(temp, d.RoundList[i:end])
		}
		templength := len(temp)
		for i := 0; i < templength; i++ {
			rt += "|巡數|First|Second|俘獲"
		}

		rt += "|\n"
		for i := 0; i < templength; i++ {
			rt += "|---|---|---|---"
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
		rt += "|巡數|First|Second|俘獲|\n"
		rt += "|---|---|---|---|\n"
		for i := 0; i < len(d.RoundList); i++ {
			rt += d.RoundList[i].Format() + "|\n"
		}
	}
	rt += "\n\n</center>\n"
	return rt
}
