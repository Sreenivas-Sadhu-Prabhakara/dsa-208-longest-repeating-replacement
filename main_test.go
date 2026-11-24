package main

import "testing"

func TestCharacterReplacement(t *testing.T) {
	cases := []struct {
		s    string
		k    int
		want int
	}{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
		{"AAAA", 0, 4},
		{"ABCD", 0, 1},
		{"ABCD", 3, 4},
		{"", 2, 0},
		{"B", 0, 1},
		{"ABBB", 2, 4},
	}
	for _, c := range cases {
		if got := CharacterReplacement(c.s, c.k); got != c.want {
			t.Errorf("CharacterReplacement(%q,%d) = %d, want %d", c.s, c.k, got, c.want)
		}
	}
}
