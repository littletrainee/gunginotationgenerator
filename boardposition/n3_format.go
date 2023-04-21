package boardposition

// 回傳格式化後棋盤上位置物件的屬性
func (b *BoatdPosition) Format() string {
	return b.Column + "-" + b.Row + "-" + b.Dan + "-"
}
