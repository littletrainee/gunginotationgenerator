package GunGiNotationGenerator

type Position struct {
	縱 string
	橫 string
	段 string
}

func (p *Position) Format() string {
	return p.縱 + "-" + p.橫 + "-" + p.段 + "-"
}

func (p *Position) SetNil() {
	p.縱 = ""
	p.橫 = ""
	p.段 = ""
}

func (p *Position) Empty() bool {
	if p.縱 == "" && p.橫 == "" && p.段 == "" {
		return true
	}
	return false
}
