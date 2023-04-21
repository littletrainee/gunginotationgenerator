package title

import (
	"fmt"
	"strconv"

	"github.com/littletrainee/ClearConsole"
	"github.com/littletrainee/gunginotationgenerator/constant"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/gunginotationgenerator/komahandler"
)

// 設置對弈階段，並初始化雙方的駒台與未被俘獲的映射
func (t *Title) SetLevel(kh *komahandler.KomaHandler) {
	var (
		key  string
		stop string
	)
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
		t.Level = level.BEGINNER
		stop = "弓"
	case "2":
		t.Level = level.ELEMENTARY
		stop = "砲"
	case "3":
		t.Level = level.INTERMEDIATE
	case "4":
		t.Level = level.ADVANCED
	}

	var (
		repeat int
		temp   string
	)
	for _, v := range constant.ALLKOMA {
		if string(v) == stop {
			break
		}
		switch string(v) {
		case "帥", "大", "中", "砲", "筒", "謀":
			repeat = 1
		case "小", "侍", "忍", "砦", "馬", "弓":
			repeat = 2
		case "槍":
			repeat = 3
		case "兵":
			repeat = 4
		}
		for i := 0; i < repeat; i++ {
			temp += string(v)
		}
	}
	kh.Initilization(temp)
	ClearConsole.ClearConsole()
}
