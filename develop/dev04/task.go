package main

import (
	"log"
	"sort"
	"strings"
)

func main() {
	data := []string{"АМКАР", "КАРМА", "КРАМА", "МАКАР", "МАКРА", "МАРКА", "РАМКА",
		"ПЯТАК", "ПЯТКА", "ТЯПКА", "КОСАЧ", "САЧОК", "ЧАСОК", "АВТОР", "ВАРТО", "ВТОРА", "ОТВАР",
		"РВОТА", "ТАВРО", "ТОВАР", "КАЧУР", "КРАУЧ", "КРУЧА", "КУРЧА", "РУЧКА", "ЧУРКА", "АБНЯ",
		"БАНЯ", "БАЯН", "КОРТ", "КРОТ", "ТРОК", "КОТ", "КТО", "ОТК", "ТОК"}
	res := SearchAnagramms(data)
	log.Print(res)
}

func SearchAnagramms(data []string) map[string][]string {
	res := make(map[string][]string)
	var flag bool
	for ind, val := range data {
		data[ind] = strings.ToLower(val)
	}

	for _, str := range data {
		for key := range res {
			flag = AnagrammsIs(key, str)
			if flag {
				res[key] = append(res[key], str)
				break
			}
		}
		if len(res) == 0 || !flag {
			res[str] = []string{}
		}
	}
	res = FormatingData(res)
	return res
}

func AnagrammsIs(key, str string) bool {
	helperMap := make(map[rune]int)

	for _, val := range key {
		if _, ok := helperMap[val]; !ok {
			helperMap[val] = 0
		}
		helperMap[val]++
	}

	for _, val := range str {
		if _, ok := helperMap[val]; !ok {
			helperMap[val] = 0
		}
		helperMap[val]++
	}

	for _, val := range helperMap {
		if val == 1 {
			return false
		}
	}
	return true
}

func FormatingData(data map[string][]string) map[string][]string {
	for key, val := range data {
		if len(val) == 0 {
			delete(data, key)
			continue
		}
		data[key] = sort.StringSlice(val)
	}
	return data
}
