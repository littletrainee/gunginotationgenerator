package GunGiNotationGenerator

import (
	"fmt"
	"strconv"
	"strings"
)

func RCL(_先後手 先後手, level Level, State State) Position {
	var (
		position Position = Position{}
		key      string
		maxlayer int
		err      error
	)
	if level != 高級 {
		maxlayer = 2
	} else {
		maxlayer = 3
	}
	for {
		fmt.Print("請輸入" + 是先手還後手(_先後手) + "的縱橫段(e/E)離開本階段：")
		fmt.Scanln(&key)
		if key == "e" || key == "E" {
			position.SetNil()
			return position
		}
		_, err = strconv.ParseInt(key, 10, 64)
		if err != nil || len(key) != 3 || strings.Contains(key, "0") {
			fmt.Println("輸入錯誤請重新輸入!")
			continue
		}
		column, err := strconv.Atoi(string(key[1]))
		if State == 布陣階段 {
			if _先後手 == 先手 {
				if err != nil || column < 7 || column > 9 {
					fmt.Println("輸入錯誤請重新輸入!")
					continue
				}
			} else {
				if err != nil || column < 1 || column > 3 {
					fmt.Println("輸入錯誤請重新輸入!")
					continue
				}
			}
		} else {
			if err != nil || column < 1 || column > 9 {
				fmt.Println("輸入錯誤請重新輸入!")
			} else {
				break
			}

		}

		layer, err := strconv.Atoi(string(key[2]))
		if err != nil || layer < 1 || layer > maxlayer {
			fmt.Println("輸入錯誤請重新輸入!")
			continue
		}
		fmt.Println("輸入的縱橫段為：" + key + "\n正確請再按一次，不正確請隨意輸入")
		var check string
		fmt.Scanln(&check)
		if check != "" {
			continue
		} else {
			break
		}
	}
	position.縱 = string(key[0])
	position.橫 = string(key[1])
	position.段 = string(key[2])

	return position
}
