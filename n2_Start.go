package gunginotationgenerator

import (
	"fmt"
	"time"

	"github.com/littletrainee/ClearConsole"
)

func (g *GunGiNotationGenerator) Start() {
	// 設定黑白雙方的名稱、先Second
	g.title.SetNameAndFirst(&g.komaHolder)
	// 設定對弈的級別與駒台
	g.title.SetLevel(&g.komaHolder)

	fmt.Println("對弈條件設置結束\n布陣開始")
	fmt.Scanln()
	ClearConsole.ClearConsole()
	g.wholeGameTime.StartTime = time.Now().Format("15:04:05")

	// 布置陣地
	g.setUp.SetUpLoop(g.komaHolder, g.title.Level)

	g.wholeGameTime.SetUpEndTime = time.Now().Format("15:04:05")
	fmt.Println("布陣結束\n對弈開始")
	fmt.Scanln()
	ClearConsole.ClearConsole()

	// 對弈期間
	g.duel.DuellingLoop(g.komaHolder, g.title.Level)

	// 設置對弈結束時間
	g.wholeGameTime.EndTime = time.Now().Format("15:04:05")
	// 彙總成string陣列
	g.list = append(g.list, g.wholeGameTime.GetTimeTitle())
	g.list = append(g.list, g.title.Format(g.komaHolder))
	g.list = append(g.list, "### <center><b>布陣階段</b></center>")
	g.list = append(g.list, g.setUp.Format())
	g.list = append(g.list, "### <center><b>對弈階段</b></center>")
	g.list = append(g.list, g.duel.Format())
	SaveToMarkDownFile(g.wholeGameTime.FileName, ConvertListToSliceByte(g.list))
}
