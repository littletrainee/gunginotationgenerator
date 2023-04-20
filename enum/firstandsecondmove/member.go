package firstandsecondmove

type FirstAndSecondMove int

const (
	FIRST FirstAndSecondMove = iota
	SECOND
)

func FirstOrSecond(目標 FirstAndSecondMove) string {
	if 目標 == FIRST {
		return "先手"
	} else {
		return "後手"
	}
}
