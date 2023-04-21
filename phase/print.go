package phase

import (
	"fmt"

	"github.com/littletrainee/gunginotationgenerator/constant"
	"github.com/littletrainee/gunginotationgenerator/enum/firstandsecondmove"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/gunginotationgenerator/enum/target"
	"github.com/littletrainee/gunginotationgenerator/komahandler"
)

func Print(kH komahandler.KomaHandler, _先Second firstandsecondmove.FirstAndSecondMove, l level.Level, t target.Target) string {
	var (
		index          int
		rt             string
		ptr            string
		targetquantity int
	)

	switch l {
	case level.BEGINNER:
		ptr = constant.ALLKOMA[:30]
	case level.ELEMENTARY:
		ptr = constant.ALLKOMA[:33]
	case level.INTERMEDIATE:
	case level.ADVANCED:
		ptr = constant.ALLKOMA
	}
	for _, v := range ptr {
		if _先Second == firstandsecondmove.FIRST {
			if t == target.KOMATAI {
				targetquantity = kH.FirstKomaTai[string(v)]
			} else if t == target.UNCAPTURED {
				targetquantity = kH.FirstKomaTaiAndBoard[string(v)]
			} else if t == target.OPPONENTBOARD {
				targetquantity = kH.SecondBoard[string(v)]
			}
		} else {
			if t == target.KOMATAI {
				targetquantity = kH.SecondKomaTai[string(v)]
			} else if t == target.UNCAPTURED {
				targetquantity = kH.SecondKomaTaiAndBoard[string(v)]
			} else if t == target.OPPONENTBOARD {
				targetquantity = kH.FirstBoard[string(v)]
			}
		}

		if targetquantity > 0 {
			fmt.Printf("%d. %s ", index+1, string(v))
			index++
			rt += string(v)
		}
	}
	fmt.Println()
	return rt
}
