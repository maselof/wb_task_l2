package main

import (
	"testing"
)

func TestSearchAnagramm(t *testing.T) {
	testVal := []struct {
		name string
		in   []string
		out  map[string][]string
	}{
		{name: "one key word", in: []string{"АМКАР", "КАРМА", "КРАМА", "МАКАР", "МАКРА", "МАРКА", "РАМКА"},
			out: map[string][]string{"амкар": {"карма", "крама", "макар", "макра", "марка", "рамка"}}},
		{name: "several key words", in: []string{"кот", "кто", "отк", "ток", "Макар", "амкар", "карма", "крама", "макра", "марка", "рамка", "сачок", "косач", "часок"},
			out: map[string][]string{
				"кот":   {"кто", "отк", "ток"},
				"макар": {"амкар", "карма", "крама", "макра", "марка", "рамка"},
				"сачок": {"косач", "часок"},
			}},
		{name: "empty key words", in: []string{}, out: map[string][]string{}},
		{name: "one word in plenty", in: []string{"собака"}, out: map[string][]string{}},
	}

	for _, testVal := range testVal {
		t.Run(testVal.name, func(t *testing.T) {
			result := SearchAnagramms(testVal.in)
			for word, anagrams := range testVal.out {
				resAnagrams, ok := result[word]
				if !ok {
					t.Errorf("expected key word '%s' in result, but it have not", word)
				}

				for i, anagram := range anagrams {
					if anagram != resAnagrams[i] {
						t.Errorf("expected anagram '%s' in result, but it have not", anagram)
					}
				}
			}
		})
	}
}
