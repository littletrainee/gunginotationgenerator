package komaholder

import "image/color"

type KomaHolder struct {
	FirstKomaTai          map[string]int
	SecondKomaTai         map[string]int
	FirstKomaTaiAndBoard  map[string]int
	SecondKomaTaiAndBoard map[string]int
	FirstBoard            map[string]int
	SecondBoard           map[string]int
	FirstColor            color.Color
	SecondColor           color.Color
}
