package GunGiNotationGenerator

type 對象 int

const (
	備牌區 對象 = iota
	未俘獲
	對方的棋盤區
)

type KomaHolder struct {
	先手備牌區    map[string]int
	後手備牌區    map[string]int
	先手備牌與棋盤區 map[string]int
	後手備牌與棋盤區 map[string]int
	先手棋盤區    map[string]int
	後手棋盤區    map[string]int
	先手顏色     Color
	後手顏色     Color
}

func (k *KomaHolder) Initilization() {
	k.先手備牌區 = make(map[string]int)
	k.後手備牌區 = make(map[string]int)
	k.先手備牌與棋盤區 = make(map[string]int)
	k.後手備牌與棋盤區 = make(map[string]int)
	k.先手棋盤區 = make(map[string]int)
	k.後手棋盤區 = make(map[string]int)
}

func (k *KomaHolder) AddtoKomaDai(s string) {
	_, firstOK := k.先手備牌區[s]
	if firstOK {
		k.先手備牌區[s]++
		k.先手備牌與棋盤區[s]++
	} else {
		k.先手備牌區[s] = 1
		k.先手備牌與棋盤區[s] = 1
	}
	_, SecondOK := k.後手備牌區[s]
	if SecondOK {
		k.後手備牌區[s]++
		k.後手備牌與棋盤區[s]++
	} else {
		k.後手備牌區[s] = 1
		k.後手備牌與棋盤區[s] = 1
	}
}

func (k *KomaHolder) SetColor(color Color) {
	if color == 黑 {
		k.先手顏色 = 黑
		k.後手顏色 = 白
	} else {
		k.先手顏色 = 白
		k.後手顏色 = 黑
	}
}
