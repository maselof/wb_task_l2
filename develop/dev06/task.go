package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	arg := os.Args

	lines := ReadDataFromTerminal()
	flag := ParseFlag(arg[1:])
	log.Print(Cut(strings.Fields(lines), flag))
}

func ReadDataFromTerminal() string {
	sc := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0, 10)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines[0]
}

func ParseFlag(flags []string) map[string]string {
	mapFlags := make(map[string]string)
	for ind, flagVal := range flags {
		if flagVal == "-d" {
			mapFlags[flagVal] = flags[ind+1]
		}
		if flagVal == "-f" {
			mapFlags[flagVal] = flags[ind+1]
		}
	}
	return mapFlags
}

func Cut(data []string, flag map[string]string) string {
	res := make([]string, 0)
	if _, ok := flag["-f"]; ok {
		getIndex := make([]int, 0)
		runeIndex := []rune(flag["-f"])
		leftIndex := 0
		flagPeriod := false
		for ind, val := range runeIndex {
			if unicode.IsDigit(val) && !flagPeriod {
				getIndex = append(getIndex, int(val-'0'))
				leftIndex = int(val - '0')
			}
			if val == '-' {
				flagPeriod = true
				if ind == len(runeIndex)-1 {
					for i := leftIndex; i < len(data); i++ {
						res = append(res, data[i])
					}
					break
				}
			}

			if flagPeriod && unicode.IsDigit(val) {
				if int(val-'0') > len(data)-1 {
					log.Fatal("incorrect period")
				}
				log.Print(leftIndex)
				for i := leftIndex; i <= int(val-'0'); i++ {
					res = append(res, data[i])
				}
				break
			}
		}
		if !flagPeriod {
			for _, valIndex := range getIndex {
				res = append(res, data[valIndex])
			}
		}
	}
	helperRes := make([]string, 0)
	if del, ok := flag["-d"]; ok {
		if len(res) != 0 {
			for _, val := range res {
				helperRes = append(helperRes, val, del)
			}
		} else {
			for _, val := range data {
				helperRes = append(helperRes, val, del)
			}
		}
	}
	strRes := ""
	if len(helperRes) == 0 {
		strRes = strings.Join(res, "")
		return strRes
	}
	strRes = strings.Join(helperRes, "")
	return strRes
}
