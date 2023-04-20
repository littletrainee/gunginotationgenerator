package arrangementphase

import "github.com/littletrainee/gunginotationgenerator/arrangementround"

// 布置陣地期間的物件
type SetUp struct {
	roundList   []arrangementround.SetUpRound
	firstIsEnd  bool
	secondIsEnd bool
}
