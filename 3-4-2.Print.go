package GunGiNotationGenerator

import "fmt"

func Print(komaHolder KomaHolder, _先後手 先後手, level Level, _對象 對象) string {
	var (
		index          int
		rt             string
		ptr            string
		targetquantity int
	)

	switch level {
	case 入門:
		ptr = 入門駒
	case 初級:
		ptr = 初級駒
	case 中級:
	case 高級:
		ptr = 所有駒
	}
	for _, v := range ptr {
		if _先後手 == 先手 {
			if _對象 == 備牌區 {
				targetquantity = komaHolder.先手備牌區[string(v)]
			} else if _對象 == 未俘獲 {
				targetquantity = komaHolder.先手備牌與棋盤區[string(v)]
			} else if _對象 == 對方的棋盤區 {
				targetquantity = komaHolder.後手棋盤區[string(v)]
			}
		} else {
			if _對象 == 備牌區 {
				targetquantity = komaHolder.後手備牌區[string(v)]
			} else if _對象 == 未俘獲 {
				targetquantity = komaHolder.後手備牌與棋盤區[string(v)]
			} else if _對象 == 對方的棋盤區 {
				targetquantity = komaHolder.先手棋盤區[string(v)]
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
