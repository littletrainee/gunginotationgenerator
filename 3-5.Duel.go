package GunGiNotationGenerator

import (
	"fmt"
	"strconv"

	"github.com/littletrainee/ClearConsole"
)

// 對弈期間的物件
type Duel struct {
	RoundList []DuelRound
	Captured  []PairCapture
	End       bool
}

// 每一巡
func (d *Duel) 本巡(komaHolder KomaHolder, index int, level Level) {
	var (
		first         string
		second        string
		firstcapture  Capture
		secondcapture Capture
	)
	first, firstcapture = d.動作(komaHolder, 先手, level)
	if first == "まげました" {
		d.End = true
		return
	}
	fmt.Println()
	second, secondcapture = d.動作(komaHolder, 後手, level)

	if second == "まげました" {
		d.End = true
		return
	}
	d.RoundList = append(d.RoundList, DuelRound{index, first, second, PairCapture{firstcapture, secondcapture}})
}

// 對弈期間的迴圈
func (d *Duel) DuelLoop(komaHolder KomaHolder, level Level) {
	var round int = 1
	for !d.End {
		fmt.Printf("第 %d 巡\n\n", round)
		d.本巡(komaHolder, round, level)
		if d.End {
			break
		} else {
			round++
		}
		ClearConsole.ClearConsole()
	}
}

func (d *Duel) 動作(komaHolder KomaHolder, _先後手 先後手, 階級 Level) (string, Capture) {
	var (
		位置      Position
		駒       string
		capture Capture
		Is新     string
	)
ReEnterRCL:
	位置 = RCL(_先後手, 階級, 對弈階段)
	if 位置.Empty() {
		return "まげました", capture
	}

	for {
		var 確認列表 = Print(komaHolder, _先後手, 階級, 未俘獲)
		if 階級 == 初級 || 階級 == 入門 {
			fmt.Println("注意：帥不可以堆疊在其他棋子上方！")
		}

		fmt.Print("請輸入選擇的駒或返回(b/B)重新選擇位置：")
		fmt.Scanln(&駒)
		if 駒 == "b" || 駒 == "B" {
			goto ReEnterRCL
		}
		選擇的數字, 錯誤 := strconv.Atoi(駒)
		if 錯誤 != nil {
			fmt.Println("輸入錯誤請重新輸入!")
			continue
		}
		駒 = getstring(確認列表, 選擇的數字-1)
		// 確認駒台是否有目標駒
		Is新 = 新(komaHolder, 駒, _先後手)
		// 有目標駒的話執行新的動作，其他的都是移駒
		if Is新 == "新" {
			FromKomaDaiToBoard(komaHolder, 駒, _先後手)
		} else {
			capture = 移駒(komaHolder, _先後手, 階級)
		}
		break

	}

	return 位置.Format() + 駒 + Is新, capture
}

// 確認是否從備排區移動到棋盤
func 新(komaHolder KomaHolder, koma string, _先後手 先後手) string {
	var key string
	if _先後手 == 先手 {
		if komaHolder.先手備牌區[koma] > 0 {
			for {
				fmt.Print(`是否要為"新加入"?(y/n)`)
				fmt.Scanln(&key)
				if key != "y" && key != "n" {
					fmt.Println("輸入錯誤請重新輸入!")
					continue
				} else {
					break
				}
			}
			if key == "y" {
				return "新"
			}
		}
	} else {
		if komaHolder.後手備牌區[koma] > 0 {
			for {
				fmt.Print(`是否要為"新加入"?(y/n)`)
				fmt.Scanln(&key)
				if key != "y" && key != "n" {
					fmt.Println("輸入錯誤請重新輸入!")
					continue
				} else {
					break
				}
			}
			if key == "y" {
				return "新"
			}
		}
	}
	return ""
}

// 移居後確認是控制還是俘獲
func 移駒(komaHolder KomaHolder, _先後手 先後手, level Level) Capture {
	var (
		key     string
		capture Capture = Capture{}
		num     int64
		err     error
	)
	for {
		fmt.Print("是否俘獲對方的駒?(y/n))")
		fmt.Scanln(&key)
		if key != "y" && key != "n" {
			fmt.Println("輸入錯誤請重新輸入!")
			continue
		} else {
			break
		}
	}
	if key == "y" {
		capture.誰俘獲 = _先後手

		for {
			fmt.Print("數量是幾個?")
			fmt.Scanln(&key)
			num, err = strconv.ParseInt(key, 10, 64)
			if err != nil || num < 1 || num > 3 {
				fmt.Println("輸入錯誤請重新輸入!")
				continue
			} else {
				break
			}
		}

		for i := 0; i < int(num); i++ {
			var (
				可能的選擇 string = Print(komaHolder, _先後手, level, 對方的棋盤區)
				選擇的數字 int
				錯誤    error
			)
			for {
				fmt.Print("輸入選擇的駒：")
				fmt.Scanln(&key)
				選擇的數字, 錯誤 = strconv.Atoi(key)
				if 錯誤 != nil || 選擇的數字 > len(可能的選擇) || 選擇的數字 < 1 {
					fmt.Println("輸入錯誤請重新輸入!")
					continue
				} else {
					break
				}
			}
			key = getstring(可能的選擇, 選擇的數字-1)
			FromBoardToOff(komaHolder, key, _先後手)
			capture.駒 += key
		}
	}
	return capture
}

// 移除對方的駒
func FromBoardToOff(komaHolder KomaHolder, koma string, _先後手 先後手) {
	if _先後手 == 後手 {
		komaHolder.先手棋盤區[koma]--
	} else {
		komaHolder.後手棋盤區[koma]--
	}
}

func (d *Duel) Format() string {
	var rt string = "<center>\n\n"
	if len(d.RoundList) > MAXROW {
		var temp [][]DuelRound
		//分割為每欄10行
		for i := 0; i < len(d.RoundList); i += 10 {
			end := i + 10
			if end > len(d.RoundList) {
				end = len(d.RoundList)
			}
			temp = append(temp, d.RoundList[i:end])
		}
		templength := len(temp)
		for i := 0; i < templength; i++ {
			rt += "|巡數|先手|後手|俘獲"
		}

		rt += "|\n"
		for i := 0; i < templength; i++ {
			rt += "|---|---|---|---"
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
		rt += "|巡數|先手|後手|俘獲|\n"
		rt += "|---|---|---|---|\n"
		for i := 0; i < len(d.RoundList); i++ {
			rt += d.RoundList[i].Format() + "|\n"
		}
	}
	rt += "\n\n</center>\n"
	return rt
}
