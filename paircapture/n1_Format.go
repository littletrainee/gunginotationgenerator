package paircapture

import "github.com/littletrainee/gunginotationgenerator/capture"

// 回傳格式化後先後手俘獲組物件的屬性
func (p *PairCapture) Format() string {
	var rt string
	if p.First != (capture.Capture{}) {
		rt += p.First.Format()
	}
	if rt != "" && p.Second != (capture.Capture{}) {
		rt += ", " + p.Second.Format()
	} else if p.Second != (capture.Capture{}) {
		rt += p.Second.Format()
	}
	return rt
}
