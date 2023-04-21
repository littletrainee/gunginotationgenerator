package phase

func ProbablyKoma(source string, selecttarget int) string {
	for i, v := range source {
		if i == selecttarget*3 {
			return string(v)
		}
	}
	return ""
}
