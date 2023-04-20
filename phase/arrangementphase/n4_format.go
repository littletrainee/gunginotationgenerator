package arrangementphase

import (
	"github.com/littletrainee/gunginotationgenerator/arrangementround"
	"github.com/littletrainee/gunginotationgenerator/constant"
)

func (s *SetUp) Format() string {
	var rt string = "<center>\n\n"

	if len(s.roundList) > constant.MAXROW {
		var temp [][]arrangementround.SetUpRound
		for i := 0; i < len(s.roundList); i += 10 {
			end := i + 10
			if end > len(s.roundList) {
				end = len((s.roundList))
			}
			temp = append(temp, s.roundList[i:end])
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
		for i := 0; i < len(s.roundList); i++ {
			rt += s.roundList[i].Format() + "|\n"
		}
	}
	rt += "\n\n</center>\n"
	return rt
}
