package arrangementphase

import (
	"fmt"

	"github.com/littletrainee/gunginotationgenerator/arrangementround"
	"github.com/littletrainee/gunginotationgenerator/enum/firstandsecondmove"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/gunginotationgenerator/komaholder"
)

// 每一巡
func (s *SetUp) turn(kh komaholder.KomaHolder, index int, l level.Level) {
	var first string
	var second string
	if !s.firstIsEnd {
		first = s.movements(kh, firstandsecondmove.FIRST, l)
		if first == "済み" {
			s.firstIsEnd = true
		}
	} else {
		first = ""
	}
	fmt.Println()
	if !s.secondIsEnd {
		second = s.movements(kh, firstandsecondmove.SECOND, l)
		if second == "済み" {
			s.secondIsEnd = true
		}
	} else {
		first = ""
	}
	s.roundList = append(s.roundList, arrangementround.SetUpRound{Order: index, First: first, Second: second})
}
