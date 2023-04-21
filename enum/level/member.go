// 階級列舉包
package level

// 階級
type Level int

const (
	BEGINNER Level = iota + 1
	ELEMENTARY
	INTERMEDIATE
	ADVANCED
)

// 回傳階級名稱
func GetLevelName(level Level) string {
	switch level {
	case BEGINNER:
		return "入門"
	case ELEMENTARY:
		return "初級"
	case INTERMEDIATE:
		return "中級"
	case ADVANCED:
		return "高級"
	default:
		return ""
	}
}
