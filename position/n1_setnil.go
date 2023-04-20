package position

func (p *Position) SetNil() {
	p.Column = ""
	p.Row = ""
	p.Dan = ""
}
