package GunGiNotationGenerator

import "time"

func (g *GunGiNotationGenerator) Initilization() {
	g.wholeGameTime = WholeGameTime{
		Date:     time.Now().Format("2006年01月02日"),
		FileName: time.Now().Format("20060102150405") + ".md"}
	g.title = Title{}
	g.komaHolder = KomaHolder{}
	g.setUp = SetUp{}
	g.duel = Duel{}
	g.komaHolder.Initilization()
}
