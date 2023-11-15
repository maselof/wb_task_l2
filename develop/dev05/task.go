package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	arg := os.Args

	flags, word := ParseFlags(arg[1:])

	lines := ReadDataFromTerminal()

	log.Print(Grep(lines, word, flags))
}

func ReadDataFromTerminal() []string {
	sc := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0, 10)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func ParseFlags(flags []string) (map[rune]int, string) {
	res := make(map[rune]int, len(flags))
	arg := make([]string, 5)
	for _, val := range flags {
		flagRune := []rune(val)
		if []rune(val)[0] == '-' {
			if len(flagRune) == 4 {
				res[flagRune[1]] = int(flagRune[3] - '0')
			} else {
				res[flagRune[1]] = 0
			}
		} else {
			arg = append(arg, val)
		}
	}
	return res, strings.Join(arg, " ")
}

func Grep(data []string, arg string, flags map[rune]int) []string {
	resHelper := make([]int, 0, len(data))
	leftIndex := 0

	arg = strings.TrimSpace(arg)

	for ind, val := range data {
		dataStrSlice := strings.Fields(val)
		for indHelper, valHelper := range dataStrSlice[leftIndex:] {
			if valHelper == arg {
				resHelper = append(resHelper, ind)
				leftIndex = indHelper
			}
		}
	}

	res := make([]string, len(resHelper))
	flag := ' '
	for k := range flags {
		flag = k
	}
	switch flag {
	case ' ':
		for _, val := range resHelper {
			res = append(res, data[val])
		}
	case 'A':
		for _, val := range resHelper {
			if len(data) > flags['A']+val {
				res = append(res, data[val:val+flags['A']+1]...)
			} else {
				res = append(res, data[val:]...)
			}
		}
	case 'B':
		for _, val := range resHelper {
			if 0 < val-flags['B'] {
				res = append(res, data[val-flags['B']:val+1]...)
			} else {
				res = append(res, data[:val+1]...)
			}
		}
	case 'C':
		for _, val := range resHelper {
			if 0 < val-flags['C'] {
				res = append(res, data[val-flags['C']:val]...)
			} else {
				res = append(res, data[:val+1]...)
			}
			if len(data) > flags['C']+val {
				res = append(res, data[val:val+flags['C']+1]...)
			} else {
				res = append(res, data[val:]...)
			}
		}
	case 'c':
		res = append(res, strconv.Itoa(len(resHelper)))
	case 'i':
		for ind, val := range data {
			dataStrSlice := strings.Fields(val)
			for indHelper, valHelper := range dataStrSlice[leftIndex:] {
				if strings.EqualFold(valHelper, arg) {
					resHelper = append(resHelper, ind)
					leftIndex = indHelper
				}
			}
		}
		for _, val := range resHelper {
			res = append(res, data[val])
		}
	case 'v':
		helperMap := make(map[int]int)
		for _, val := range resHelper {
			helperMap[val] = 1
		}

		for ind, val := range data {
			if _, ok := helperMap[ind]; !ok {
				res = append(res, val)
			}
		}
	case 'F':
		for _, val := range data {
			val = strings.TrimSpace(val)
			if strings.EqualFold(val, arg) {
				res = append(res, val)
			}
		}
	case 'n':
		for _, val := range resHelper {
			res = append(res, strconv.Itoa(val))
		}
	}

	if len(res) == 0 {
		log.Printf("Not found word %s", arg)
		return nil
	} else {
		return WriteResult(res)
	}
}

func WriteResult(resHelper []string) []string {
	res := make([]string, 0)
	for _, val := range resHelper {
		if strings.TrimSpace(val) != "" && strings.TrimSpace(val) != " " {
			res = append(res, strings.TrimSpace(val))
		}
	}
	return res
}
