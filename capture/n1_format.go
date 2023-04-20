package capture

import "github.com/littletrainee/gunginotationgenerator/enum/firstandsecondmove"

func (c *Capture) Format() string {
	if c.CapturedBy == firstandsecondmove.FIRST {
		return "First" + "-" + c.Koma
	} else {
		return "Second" + "-" + c.Koma
	}
}
