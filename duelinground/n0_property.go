package duelinground

import "github.com/littletrainee/gunginotationgenerator/paircapture"

type DuelRound struct {
	Order       int
	First       string
	Second      string
	PairCapture paircapture.PairCapture
}
