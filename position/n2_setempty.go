package position

func (p *Position) Empty() bool {
	if p.Column == "" && p.Row == "" && p.Dan == "" {
		return true
	}
	return false
}
