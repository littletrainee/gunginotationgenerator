package GunGiNotationGenerator

import (
	"fmt"
	"os"
)

// filename is Year+Month+Day+Hour+Minute+second, everyone can read, write and explore
func SaveToMarkDownFile(filename string, s []byte) {
	err := os.WriteFile(filename, s, 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ConvertListToArraryBute(source []string) []byte {
	var temp string
	for i, v := range source {
		if i != 0 {
			temp += "\n"
		}
		temp += v
	}
	return []byte(temp)
}
