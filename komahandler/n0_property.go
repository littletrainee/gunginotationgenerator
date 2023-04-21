// 駒管理者包
package komahandler

import "image/color"

// 駒管理者物件
type KomaHandler struct {
	FirstKomaTai          map[string]int
	SecondKomaTai         map[string]int
	FirstKomaTaiAndBoard  map[string]int
	SecondKomaTaiAndBoard map[string]int
	FirstBoard            map[string]int
	SecondBoard           map[string]int
	FirstColor            color.Color
	SecondColor           color.Color
}
