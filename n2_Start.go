package gunginotationgenerator

import (
	"fmt"
	"time"

	"github.com/littletrainee/ClearConsole"
)

// 開始運行
func (g *GunGiNotationGenerator) Start() {
	// 設定黑白雙方的名稱、先Second
	g.title.SetNameAndFirst(&g.komaHolder)
	// 設定對弈的級別與駒台
	g.title.SetLevel(&g.komaHolder)

	fmt.Println("對弈條件設置結束\n布陣開始")
	fmt.Scanln()
	ClearConsole.ClearConsole()
	g.timeRecord.StartTime = time.Now().Format("15:04:05")

	// 布置陣地
	g.arrangementPhase.Loop(g.komaHolder, g.title.Level)

	g.timeRecord.SetUpEndTime = time.Now().Format("15:04:05")
	fmt.Println("布陣結束\n對弈開始")
	fmt.Scanln()
	ClearConsole.ClearConsole()

	// 對弈期間
	g.duelingPhase.Loop(g.komaHolder, g.title.Level)

	// 設置對弈結束時間
	g.timeRecord.EndTime = time.Now().Format("15:04:05")
	// 彙總成string陣列
	g.stringToOutput = append(g.stringToOutput, g.timeRecord.GetTimeTitle())
	g.stringToOutput = append(g.stringToOutput, g.title.Format(g.komaHolder))
	g.stringToOutput = append(g.stringToOutput, "### <center><b>布陣階段</b></center>")
	g.stringToOutput = append(g.stringToOutput, g.arrangementPhase.Format())
	g.stringToOutput = append(g.stringToOutput, "### <center><b>對弈階段</b></center>")
	g.stringToOutput = append(g.stringToOutput, g.duelingPhase.Format())
	SaveToMarkDownFile(g.timeRecord.FileName, ConvertToSliceByte(g.stringToOutput))
}
