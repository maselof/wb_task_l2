package main

import (
	"log"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	TestData := []struct {
		name []string
		in   []string
		out  []string
	}{
		{
			name: []string{"no flags", ""},
			in:   []string{"1", "4", "5", "2", "3"},
			out:  []string{"1", "2", "3", "4", "5"},
		},
		{
			name: []string{"flags", "-u"},
			in:   []string{"1", "4", "5", "5", "3"},
			out:  []string{"1", "3", "4", "5"},
		},
		{
			name: []string{"flags", "-r"},
			in:   []string{"1", "4", "5", "2", "3"},
			out:  []string{"5", "4", "3", "2", "1"},
		},
		{
			name: []string{"flags", "-k=2"},
			in:   []string{"1 5", "4 2", "5 1", "2 4", "3 3"},
			out:  []string{"5 1", "4 2", "3 3", "2 4", "1 5"},
		},
	}

	for _, testVal := range TestData {
		t.Run(strings.Join(testVal.name, " "), func(t *testing.T) {
			result, err := SortDataFromFile(testVal.in, testVal.name[1])
			if err != nil {
				log.Fatal(err)
			}
			log.Print(result, testVal.out)
			for i := 0; i < len(result); i++ {
				if result[i] != testVal.out[i] {
					log.Fatal("Incorrect result")
				}
			}
		})
	}
}
