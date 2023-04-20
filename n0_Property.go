package gunginotationgenerator

import (
	"github.com/littletrainee/gunginotationgenerator/fulltime"
	"github.com/littletrainee/gunginotationgenerator/komaholder"
	"github.com/littletrainee/gunginotationgenerator/phase/arrangementphase"
	"github.com/littletrainee/gunginotationgenerator/phase/duelingphase"
	"github.com/littletrainee/gunginotationgenerator/title"
)

type GunGiNotationGenerator struct {
	list          []string
	wholeGameTime fulltime.FullTime
	title         title.Title
	komaHolder    komaholder.KomaHolder
	setUp         arrangementphase.SetUp
	duel          duelingphase.Duel
}
