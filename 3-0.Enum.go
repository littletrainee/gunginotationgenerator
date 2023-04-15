package GunGiNotationGenerator

type Level int

const (
	入門 Level = iota + 1
	初級
	中級
	高級
)

type Color int

const (
	白 Color = iota
	黑
)

type 先後手 int

const (
	先手 先後手 = iota
	後手
)

func 獲得階級名稱(level Level) string {
	switch level {
	case 入門:
		return "入門"
	case 初級:
		return "初級"
	case 中級:
		return "中級"
	case 高級:
		return "高級"
	default:
		return ""
	}
}

func 是先手還後手(目標 先後手) string {
	if 目標 == 先手 {
		return "先手"
	} else {
		return "後手"
	}
}

type State int

const (
	布陣階段 State = iota
	對弈階段
)

const 所有駒 string = "帥大中小侍忍槍砦馬兵弓砲筒謀"
const 初級駒 string = "帥大中小侍忍槍砦馬兵弓"
const 入門駒 string = "帥大中小侍忍槍砦馬兵"
const MAXROW int = 10
