package GunGiNotationGenerator

type Capture struct {
	誰俘獲 先後手
	駒   string
}

func (c *Capture) Format() string {
	if c.誰俘獲 == 先手 {
		return "先手" + "-" + c.駒
	} else {
		return "後手" + "-" + c.駒
	}
}

type PairCapture struct {
	先手 Capture
	後手 Capture
}

func (p *PairCapture) Format() string {
	var rt string
	if p.先手 != (Capture{}) {
		rt += p.先手.Format()
	}
	if rt != "" && p.後手 != (Capture{}) {
		rt += ", " + p.後手.Format()
	} else if p.後手 != (Capture{}) {
		rt += p.後手.Format()
	}
	return rt
}
