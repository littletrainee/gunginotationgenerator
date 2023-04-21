package capture

import "github.com/littletrainee/gunginotationgenerator/enum/firstandsecondmove"

// 回傳格式化後俘獲物件的屬性
func (c *Capture) Format() string {
	if c.CapturedBy == firstandsecondmove.FIRST {
		return "First" + "-" + c.Koma
	} else {
		return "Second" + "-" + c.Koma
	}
}
