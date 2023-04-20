package gunginotationgenerator

import (
	"time"

	"github.com/littletrainee/gunginotationgenerator/fulltime"
	"github.com/littletrainee/gunginotationgenerator/komaholder"
	"github.com/littletrainee/gunginotationgenerator/phase/arrangementphase"
	"github.com/littletrainee/gunginotationgenerator/phase/duelingphase"
	"github.com/littletrainee/gunginotationgenerator/title"
)

func (g *GunGiNotationGenerator) Initilization() {
	g.wholeGameTime = fulltime.FullTime{
		Date:     time.Now().Format("2006年01月02日"),
		FileName: time.Now().Format("20060102150405") + ".md"}
	g.title = title.Title{}
	g.komaHolder = komaholder.KomaHolder{}
	g.setUp = arrangementphase.SetUp{}
	g.duel = duelingphase.Duel{}
	g.komaHolder.Initilization()
}
