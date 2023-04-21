package komahandler

// 初始化
func (k *KomaHandler) Initilization(s string) {
	k.FirstKomaTai = make(map[string]int)
	k.SecondKomaTai = make(map[string]int)
	k.FirstKomaTaiAndBoard = make(map[string]int)
	k.SecondKomaTaiAndBoard = make(map[string]int)
	k.FirstBoard = make(map[string]int)
	k.SecondBoard = make(map[string]int)

	for _, v := range s {
		_, firstOK := k.FirstKomaTai[string(v)]
		if firstOK {
			k.FirstKomaTai[string(v)]++
			k.FirstKomaTaiAndBoard[string(v)]++
		} else {
			k.FirstKomaTai[string(v)] = 1
			k.FirstKomaTaiAndBoard[string(v)] = 1
		}
		_, SecondOK := k.SecondKomaTai[string(v)]
		if SecondOK {
			k.SecondKomaTai[string(v)]++
			k.SecondKomaTaiAndBoard[string(v)]++
		} else {
			k.SecondKomaTai[string(v)] = 1
			k.SecondKomaTaiAndBoard[string(v)] = 1
		}

	}
}
