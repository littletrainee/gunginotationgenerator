package duelinground

import "strconv"

// 回傳格式化後對弈巡的屬性
func (r *DuelingRound) Format() string {
	num := strconv.Itoa(r.Order)
	return "|" + num + "|" + r.First + "|" + r.Second + "|" + r.PairCapture.Format()
}
