package komaholder

func (k *KomaHolder) Initilization() {
	k.FirstKomaTai = make(map[string]int)
	k.SecondKomaTai = make(map[string]int)
	k.FirstKomaTaiAndBoard = make(map[string]int)
	k.SecondKomaTaiAndBoard = make(map[string]int)
	k.FirstBoard = make(map[string]int)
	k.SecondBoard = make(map[string]int)
}
