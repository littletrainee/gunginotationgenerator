package level

type Level int

const (
	BEGINNER Level = iota + 1
	ELEMENTARY
	INTERMEDIATE
	ADVANCED
)

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
