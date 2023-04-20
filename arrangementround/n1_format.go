package arrangementround

import "strconv"

func (r *SetUpRound) Format() string {
	num := strconv.Itoa(r.Order)
	return "|" + num + "|" + r.First + "|" + r.Second
}
