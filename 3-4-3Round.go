package GunGiNotationGenerator

import "strconv"

type SetUpRound struct {
	Order  int
	First  string
	Second string
}

func (r *SetUpRound) Format() string {
	num := strconv.Itoa(r.Order)
	return "|" + num + "|" + r.First + "|" + r.Second
}

type DuelRound struct {
	Order       int
	First       string
	Second      string
	PairCapture PairCapture
}

func (r *DuelRound) Format() string {
	num := strconv.Itoa(r.Order)
	return "|" + num + "|" + r.First + "|" + r.Second + "|" + r.PairCapture.Format()
}
