package duelinground

import "strconv"

func (r *DuelRound) Format() string {
	num := strconv.Itoa(r.Order)
	return "|" + num + "|" + r.First + "|" + r.Second + "|" + r.PairCapture.Format()
}
