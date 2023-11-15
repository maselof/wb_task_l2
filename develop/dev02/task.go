package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"
)

// Задача на распоковку строки

func main() {
	str := ""
	log.Print(unpackString(str))
}

func unpackString(str string) (string, error) {
	runeStr := []rune(str)
	if len(runeStr) == 0 {
		return "", nil
	}
	if unicode.IsDigit(runeStr[0]) {
		return "", fmt.Errorf("incorrect string")
	}

	var b strings.Builder

	for i := 0; i < len(str); i++ {
		if unicode.IsLetter(runeStr[i]) {
			fmt.Fprintf(&b, "%s", string(runeStr[i]))
		}
		if unicode.IsNumber(runeStr[i]) {
			for j := 0; j < int(runeStr[i]-'0')-1; j++ {
				fmt.Fprintf(&b, "%s", string(runeStr[i-1]))
			}
		}
	}
	str = b.String()
	return str, nil
}
