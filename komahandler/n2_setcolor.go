package komahandler

import "image/color"

//設定先後手顏色
func (k *KomaHandler) SetColor(Color color.Color) {
	if Color == color.Black {
		k.FirstColor = color.Black
		k.SecondColor = color.White
	} else {
		k.FirstColor = color.White
		k.SecondColor = color.Black
	}
}
