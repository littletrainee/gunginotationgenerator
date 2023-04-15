package GunGiNotationGenerator

import (
	"fmt"
	"strconv"

	"github.com/littletrainee/ClearConsole"
)

type Title struct {
	h3mark     string
	black      string
	white      string
	blackchess string
	whitechess string
	Level      Level
}

func (t *Title) SetNameAndFirst(komaHolder *KomaHolder) {
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
		komaHolder.SetColor(黑)
	} else if key == "w" {
		komaHolder.SetColor(白)
	}
	fmt.Println()
}

func (t *Title) SetLevel(komaHolder *KomaHolder) {
	var (
		key  string
		stop string
	)
	// t.KomaInUse = make(map[string]int)
	for {
		fmt.Print("1.入門 2.初級 3.中級 4.上級\n對弈級別：")
		fmt.Scanln(&key)
		num, err := strconv.Atoi(key)
		if err != nil || num < 1 || num > 4 {
			fmt.Print("輸入錯誤請重新輸入！")
			continue
		} else {
			break
		}
	}
	switch key {
	case "1":
		t.Level = 入門
		stop = "弓"
	case "2":
		t.Level = 初級
		stop = "砲"
	case "3":
		t.Level = 中級
	case "4":
		t.Level = 高級
	}

	var repeat int
	for _, v := range 所有駒 {
		if string(v) == stop {
			break
		}
		switch string(v) {
		case "帥", "大", "中", "砲", "筒", "謀":
			repeat = 1
			// t.KomaInUse = append(t.KomaInUse, Koma{Name: string(v)})
		case "小", "侍", "忍", "砦", "馬", "弓":
			repeat = 2
		case "槍":
			repeat = 3
		case "兵":
			repeat = 4
		}
		for i := 0; i < repeat; i++ {
			komaHolder.AddtoKomaDai(string(v))
		}
	}
	ClearConsole.ClearConsole()
}

func (t *Title) Format(komaHolder KomaHolder) string {
	var rt string = t.h3mark
	if komaHolder.先手顏色 == 黑 {
		rt += t.blackchess + "黑" + "</span>" + "棋方：" + t.black + "<br/>" +
			t.whitechess + "白" + "</span>" + "棋方：" + t.white + "<br/>"
	} else {
		rt += t.whitechess + "白" + "</span>" + "棋方：" + t.white + "<br/>" +
			t.blackchess + "黑" + "</span>" + "棋方：" + t.black + "<br/>"
	}
	rt += "對弈等級：" + 獲得階級名稱(t.Level) + "\n\n"
	return rt
}
