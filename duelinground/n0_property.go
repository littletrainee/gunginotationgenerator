// 對弈巡包
package duelinground

import "github.com/littletrainee/gunginotationgenerator/paircapture"

// 對弈巡物件
type DuelingRound struct {
	Order       int
	First       string
	Second      string
	PairCapture paircapture.PairCapture
}
