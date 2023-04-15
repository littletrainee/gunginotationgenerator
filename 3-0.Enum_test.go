package GunGiNotationGenerator

import (
	"fmt"
	"testing"
)

func TestGetLevelName(t *testing.T) {
	fmt.Println(獲得階級名稱(入門))
}

func TestPrintlen(t *testing.T) {
	for _, v := range 所有駒 {
		fmt.Println(string(v))
	}
}
