package gunginotationgenerator

import (
	"fmt"
	"os"
)

// 將結果寫入到MarkDown文件裡面
func SaveToMarkDownFile(filename string, s []byte) {
	// filename is Year+Month+Day+Hour+Minute+second, everyone can read, write and explore
	err := os.WriteFile(filename, s, 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// 將字串轉為位元切片
func ConvertToSliceByte(source []string) []byte {
	var temp string
	for i, v := range source {
		if i != 0 {
			temp += "\n"
		}
		temp += v
	}
	return []byte(temp)
}
