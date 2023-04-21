// 布陣期間包
package arrangementphase

import "github.com/littletrainee/gunginotationgenerator/arrangementround"

// 布置陣地期間物件
type ArrangementPhase struct {
	roundList   []arrangementround.ArrangementRound
	firstIsEnd  bool
	secondIsEnd bool
}
