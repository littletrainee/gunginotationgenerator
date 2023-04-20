package paircapture

import "github.com/littletrainee/gunginotationgenerator/capture"

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
