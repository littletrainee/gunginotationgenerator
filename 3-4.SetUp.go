package GunGiNotationGenerator

import (
	"fmt"
	"strconv"

	"github.com/littletrainee/ClearConsole"
)

// 布置陣地期間的物件
type SetUp struct {
	巡列表  []SetUpRound
	先手結束 bool
	後手結束 bool
}

// 布置陣地期間的迴圈
func (s *SetUp) SetUpLoop(komaHolder KomaHolder, level Level) {
	var round int = 1
	for {
		fmt.Printf("第 %d 巡\n\n", round)
		s.本巡(komaHolder, round, level)
		if s.先手結束 && s.後手結束 {
			break
		} else {
			round++
		}
		ClearConsole.ClearConsole()
	}
}

// 每一巡
func (s *SetUp) 本巡(komaHolder KomaHolder, index int, level Level) {
	var first string
	var second string
	if !s.先手結束 {
		first = s.動作(komaHolder, 先手, level)
		if first == "済み" {
			s.先手結束 = true
		}
	} else {
		first = ""
	}
	fmt.Println()
	if !s.後手結束 {
		second = s.動作(komaHolder, 後手, level)
		if second == "済み" {
			s.後手結束 = true
		}
	} else {
		first = ""
	}
	s.巡列表 = append(s.巡列表, SetUpRound{index, first, second})
}

func (s *SetUp) 動作(komaHolder KomaHolder, _先後手 先後手, 階級 Level) string {
	var (
		位置 Position = Position{}
		駒  string
	)
ReEnterRCL:
	位置 = RCL(_先後手, 階級, 布陣階段)
	if 位置.Empty() {
		return "済み"
	}
	for {
		var 可能的選擇 string = Print(komaHolder, _先後手, 階級, 備牌區)
		if 階級 == 初級 || 階級 == 入門 {
			fmt.Println("注意：帥不可以堆疊在其他棋子上方！")
		}

		fmt.Print("請輸入選擇的駒：")
		fmt.Scanln(&駒)
		if 駒 == "b" {
			goto ReEnterRCL
		}
		選擇的數字, 錯誤 := strconv.Atoi(駒)
		if 錯誤 != nil {
			fmt.Println("輸入錯誤請重新輸入!")
			continue
		}
		駒 = getstring(可能的選擇, 選擇的數字-1)
		FromKomaDaiToBoard(komaHolder, 駒, _先後手)
		break
	}

	return 位置.Format() + 駒
}

func getstring(source string, selecttarget int) string {
	for i, v := range source {
		if i == selecttarget*3 {
			return string(v)
		}
	}
	return ""
}

// 從駒台到棋盤
func FromKomaDaiToBoard(komaHolder KomaHolder, koma string, _先後手 先後手) {
	if _先後手 == 先手 {
		komaHolder.先手備牌區[koma]--
		_, 存在 := komaHolder.先手棋盤區[koma]
		if 存在 {
			komaHolder.先手棋盤區[koma]++
		} else {
			komaHolder.先手棋盤區[koma] = 1
		}
	} else {
		komaHolder.後手備牌區[koma]--
		_, 存在 := komaHolder.後手棋盤區[koma]
		if 存在 {
			komaHolder.後手棋盤區[koma]++
		} else {
			komaHolder.後手棋盤區[koma] = 1
		}
	}

}

func (s *SetUp) Format() string {
	var rt string = "<center>\n\n"

	if len(s.巡列表) > MAXROW {
		var temp [][]SetUpRound
		for i := 0; i < len(s.巡列表); i += 10 {
			end := i + 10
			if end > len(s.巡列表) {
				end = len((s.巡列表))
			}
			temp = append(temp, s.巡列表[i:end])
		}

		templength := len(temp)
		for i := 0; i < templength; i++ {
			rt += "|巡數|先手|後手"

		}
		rt += "|\n"
		lastSliceIndex := 0
		for i := 0; i < 10; i++ {
			for j := 0; j < templength; j++ {
				if len(temp[j]) == len(temp[0]) {
					rt += temp[j][i].Format()
				} else if lastSliceIndex < len(temp[len(temp)-1]) {
					rt += temp[j][lastSliceIndex].Format()
					lastSliceIndex++
				}
			}
			rt += "|\n"
		}
	} else {
		rt += "|巡數|先手|後手|\n"
		rt += "|---|---|---|\n"
		for i := 0; i < len(s.巡列表); i++ {
			rt += s.巡列表[i].Format() + "|\n"
		}
	}
	rt += "\n\n</center>\n"
	return rt
}
