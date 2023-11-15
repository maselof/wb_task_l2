package main

import (
	"log"
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	testVals := []struct {
		name string
		in   []string
		out  string
	}{
		{
			name: "flag -d ;",
			in:   []string{"a", "b", "c", "d", "e"},
			out:  "a;b;c;d;e;",
		},
		{
			name: "flag -f 1-2",
			in:   []string{"a", "b", "c", "d", "e"},
			out:  "bc",
		},
		{
			name: "flag -d ; -f 1,2,3",
			in:   []string{"a", "b", "c", "d", "e"},
			out:  "b;c;d;",
		},
	}

	for _, testVal := range testVals {
		t.Run(testVal.name, func(t *testing.T) {
			flags := ParseFlag(strings.Fields(testVal.name)[1:])
			res := Cut(testVal.in, flags)
			log.Print(res, testVal.out)
			if res != testVal.out {
				log.Fatal("Incorrect res")
			}
		})
	}
}
