// 軍儀棋棋譜產生器
package gunginotationgenerator

import (
	"github.com/littletrainee/gunginotationgenerator/komahandler"
	"github.com/littletrainee/gunginotationgenerator/phase/arrangementphase"
	"github.com/littletrainee/gunginotationgenerator/phase/duelingphase"
	"github.com/littletrainee/gunginotationgenerator/timerecord"
	"github.com/littletrainee/gunginotationgenerator/title"
)

// 軍儀棋棋譜產生器物件
type GunGiNotationGenerator struct {
	stringToOutput   []string
	timeRecord       timerecord.TimeRecord
	title            title.Title
	komaHolder       komahandler.KomaHandler
	arrangementPhase arrangementphase.ArrangementPhase
	duelingPhase     duelingphase.DuelingPhase
}
