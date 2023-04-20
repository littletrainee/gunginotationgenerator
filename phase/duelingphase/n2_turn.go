package duelingphase

import (
	"fmt"

	"github.com/littletrainee/gunginotationgenerator/capture"
	"github.com/littletrainee/gunginotationgenerator/duelinground"
	"github.com/littletrainee/gunginotationgenerator/enum/firstandsecondmove"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/gunginotationgenerator/komaholder"
	"github.com/littletrainee/gunginotationgenerator/paircapture"
)

// 每一巡
func (d *Duel) turn(komaHolder komaholder.KomaHolder, index int, level level.Level) {
	var (
		first         string
		second        string
		firstcapture  capture.Capture
		secondcapture capture.Capture
	)
	first, firstcapture = d.movements(komaHolder, firstandsecondmove.FIRST, level)
	if first == "まげました" {
		d.End = true
		return
	}
	fmt.Println()
	second, secondcapture = d.movements(komaHolder, firstandsecondmove.SECOND, level)

	if second == "まげました" {
		d.End = true
		return
	}
	d.RoundList = append(d.RoundList, duelinground.DuelRound{
		Order:  index,
		First:  first,
		Second: second,
		PairCapture: paircapture.PairCapture{
			First:  firstcapture,
			Second: secondcapture}})
}
