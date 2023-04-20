package komaholder

func (k *KomaHolder) AddtoKomaDai(s string) {
	_, firstOK := k.FirstKomaTai[s]
	if firstOK {
		k.FirstKomaTai[s]++
		k.FirstKomaTaiAndBoard[s]++
	} else {
		k.FirstKomaTai[s] = 1
		k.FirstKomaTaiAndBoard[s] = 1
	}
	_, SecondOK := k.SecondKomaTai[s]
	if SecondOK {
		k.SecondKomaTai[s]++
		k.SecondKomaTaiAndBoard[s]++
	} else {
		k.SecondKomaTai[s] = 1
		k.SecondKomaTaiAndBoard[s] = 1
	}
}
