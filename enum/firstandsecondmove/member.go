// 先後手包
package firstandsecondmove

// 先後手
type FirstAndSecondMove int

const (
	FIRST FirstAndSecondMove = iota
	SECOND
)

// 回傳先手或後手
func FirstOrSecond(目標 FirstAndSecondMove) string {
	if 目標 == FIRST {
		return "先手"
	} else {
		return "後手"
	}
}
