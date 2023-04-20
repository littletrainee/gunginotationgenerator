package komaholder

import "image/color"

func (k *KomaHolder) SetColor(Color color.Color) {
	if Color == color.Black {
		k.FirstColor = color.Black
		k.SecondColor = color.White
	} else {
		k.FirstColor = color.White
		k.SecondColor = color.Black
	}
}
