package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	// Читаем данные с файла
	data, err := ReadLinesFromFile("hello.txt")
	if err != nil {
		log.Print("Can't open file to read: ", err)
	}

	// Обрабатываем аргумент
	flags := os.Args
	flag, err := ParserFlags(flags)
	if err != nil {
		log.Fatal(err)
	}

	// Сортируем данные
	data, err = SortDataFromFile(data, flag)
	if err != nil {
		log.Fatal(err)
	}

	// Записываем данные в файл
	res, err := WriteDataToFile(data, "hello1.txt")

	// Вывод результата
	if res {
		log.Print("Data sorted and wrote successfull")
	} else {
		log.Print("Can't open file to write data: ", err)
	}
}

// Функция для обработки флагов
func ParserFlags(flags []string) (string, error) {
	flagsExample := []rune{'k', 'n', 'r', 'u'}
	correctFlag := false
	flagsRune := []rune(flags[1])
	if len(flags) == 1 {
		return "", nil
	}
	for _, val := range flagsExample {
		if val == flagsRune[1] {
			correctFlag = true
			break
		}
	}
	if correctFlag {
		return flags[1], nil
	} else {
		return "", fmt.Errorf("flag does exists")
	}
}

// Функция для считывания строк в файле
func ReadLinesFromFile(fileName string) ([]string, error) {
	res := []string{}

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		res = append(res, sc.Text())
	}

	return res, nil
}

// Функция для сортировки данных, с используемым флагом
func SortDataFromFile(data []string, flag string) ([]string, error) {
	flagRune := []rune(flag)
	switch {
	case flag == "":
		sort.Strings(data)
	case flag == "-r":
		sort.Strings(data)

		for i, j := 0, len(data)-1; i < len(data)/2; i, j = i+1, j-1 {
			data[i], data[j] = data[j], data[i]
		}
	case flag == "-u":
		helperMap := make(map[string]int)
		count := 0
		for _, val := range data {
			if _, ok := helperMap[val]; !ok {
				count++
			}
			helperMap[val]++
		}
		helperSlice := make([]string, 0, count)
		for str := range helperMap {
			helperSlice = append(helperSlice, str)
		}
		sort.Strings(helperSlice)
		data = helperSlice
	case string(flagRune[:2]) == "-k":
		log.Print(string(flagRune[3]))
		for ind, val := range data {
			helperSlice := strings.Fields(val)
			if len(helperSlice) < int(flagRune[3]-'0') {
				return nil, fmt.Errorf("invalid val")
			}
			helperSlice[0], helperSlice[int(flagRune[3]-'0')-1] = helperSlice[int(flagRune[3]-'0')-1], helperSlice[0]
			data[ind] = strings.Join(helperSlice[:], " ")
		}
		sort.Strings(data)
		for ind, val := range data {
			helperSlice := strings.Fields(val)
			helperSlice[0], helperSlice[int(flagRune[3]-'0')-1] = helperSlice[int(flagRune[3]-'0')-1], helperSlice[0]
			data[ind] = strings.Join(helperSlice[:], " ")
		}
	}
	return data, nil
}

// Запись отсортируемых данные в файл
func WriteDataToFile(data []string, fileName string) (bool, error) {
	f, err := os.Create(fileName)
	if err != nil {
		return false, err
	}
	defer f.Close()
	for _, val := range data {
		f.WriteString(val + "\n")
	}

	return true, nil
}
