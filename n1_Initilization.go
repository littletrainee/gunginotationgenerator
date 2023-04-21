package gunginotationgenerator

import (
	"time"

	"github.com/littletrainee/gunginotationgenerator/komahandler"
	"github.com/littletrainee/gunginotationgenerator/phase/arrangementphase"
	"github.com/littletrainee/gunginotationgenerator/phase/duelingphase"
	"github.com/littletrainee/gunginotationgenerator/timerecord"
	"github.com/littletrainee/gunginotationgenerator/title"
)

// 初始化
func (g *GunGiNotationGenerator) Initilization() {
	g.timeRecord = timerecord.TimeRecord{
		Date:     time.Now().Format("2006年01月02日"),
		FileName: time.Now().Format("20060102150405") + ".md"}
	g.title = title.Title{}
	g.komaHolder = komahandler.KomaHandler{}
	g.arrangementPhase = arrangementphase.ArrangementPhase{}
	g.duelingPhase = duelingphase.DuelingPhase{}
}
