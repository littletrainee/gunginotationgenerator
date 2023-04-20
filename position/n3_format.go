package position

func (p *Position) Format() string {
	return p.Column + "-" + p.Row + "-" + p.Dan + "-"
}
