package boardposition

// 是否為空值
func (b *BoatdPosition) IsEmpty() bool {
	if b.Column == "" && b.Row == "" && b.Dan == "" {
		return true
	}
	return false
}
