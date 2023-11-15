package main

import (
	"log"
	"strings"
	"testing"
)

func TestGrep(t *testing.T) {
	testVals := []struct {
		name []string
		word string
		in   []string
		out  []string
	}{
		{
			name: []string{"flag", "-A=1"},
			word: "amr",
			in:   []string{"bla", "abr amr", "cup", "adel", "achur amr"},
			out:  []string{"abr amr", "cup", "achur amr"},
		},
		{
			name: []string{"flag", "-B=1"},
			word: "amr",
			in:   []string{"bla", "abr amr", "cup", "adel", "achur amr"},
			out:  []string{"bla", "abr amr", "adel", "achur amr"},
		},
		{
			name: []string{"flag", "-C=1"},
			word: "cup",
			in:   []string{"bla", "abr amr", "cup", "adel", "achur amr"},
			out:  []string{"abr amr", "cup", "adel"},
		},
		{
			name: []string{"flag", "-c"},
			word: "amr",
			in:   []string{"bla", "abr amr", "cup", "adel", "achur amr"},
			out:  []string{"2"},
		},
		{
			name: []string{"flag", "-i"},
			word: "aMr",
			in:   []string{"bla", "abr amr", "cup", "adel", "achur amr"},
			out:  []string{"abr amr", "achur amr"},
		},
		{
			name: []string{"flag", "-v"},
			word: "amr",
			in:   []string{"bla", "abr amr", "cup", "adel", "achur amr"},
			out:  []string{"bla", "cup", "adel"},
		},
		{
			name: []string{"flag", "-F"},
			word: "abr amr",
			in:   []string{"bla", "abr amr", "cup", "adel", "achur amr"},
			out:  []string{"abr amr"},
		},
		{
			name: []string{"flag", "-n"},
			word: "amr",
			in:   []string{"bla", "abr amr", "cup", "adel", "achur amr"},
			out:  []string{"1", "4"},
		},
	}

	for _, testVal := range testVals {
		t.Run(strings.Join(testVal.name, " "), func(t *testing.T) {
			flags, _ := ParseFlags(testVal.name)
			result := Grep(testVal.in, testVal.word, flags)
			for i := 0; i < len(result); i++ {
				if result[i] != testVal.out[i] {
					log.Fatal("incorrect output")
				}
			}
		})
	}
}
