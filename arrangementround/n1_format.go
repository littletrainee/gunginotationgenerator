package arrangementround

import "strconv"

// 回傳格式化後布陣巡的屬性
func (a *ArrangementRound) Format() string {
	num := strconv.Itoa(a.Order)
	return "|" + num + "|" + a.First + "|" + a.Second
}
