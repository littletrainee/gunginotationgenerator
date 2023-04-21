package title

import (
	"fmt"
	"image/color"

	"github.com/littletrainee/ClearConsole"
	"github.com/littletrainee/gunginotationgenerator/komahandler"
)

// 設定雙方的名稱與先後手
func (t *Title) SetNameAndFirst(kh *komahandler.KomaHandler) {
	var (
		key string
	)
	t.h3mark = "### "
	t.blackchess = `<span style="background-color:black; color:white;">`
	t.whitechess = `<span style="background-color:white; color:black;">`
	ClearConsole.ClearConsole()
	fmt.Print("持白棋方名稱：")
	fmt.Scanln(&key)
	t.white = key
	fmt.Println()
	fmt.Print("持黑棋方名稱：")
	fmt.Scanln(&key)
	t.black = key
	fmt.Println()
	for {
		fmt.Print("先手為黑(b)方還是白(w)方?：")
		fmt.Scanln(&key)
		if key != "b" && key != "w" {
			fmt.Println("輸入錯誤請重新輸入選擇(b/w))!")
		} else {
			break
		}
	}

	if key == "b" {
		kh.SetColor(color.Black)
	} else if key == "w" {
		kh.SetColor(color.White)
	}
	fmt.Println()
}
