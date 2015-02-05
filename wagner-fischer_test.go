package smetrics

import "testing"

func TestWagnerFischer(t *testing.T) {
	cases := []levenshteincase{
		{"RASH", "RASH", 1, 1, 2, 0},
		{"POTATO", "POTTATO", 1, 1, 2, 1},
		{"POTTATO", "POTATO", 1, 1, 2, 1},
		{"HOUSE", "MOUSE", 1, 1, 2, 2},
		{"MOUSE", "HOUSE", 2, 2, 4, 4},
	}

	for _, c := range cases {
		if r := WagnerFischer(c.s, c.t, c.icost, c.dcost, c.scost); r != c.r {
			t.Errorf("%v, instead of %v", r, c.r)
		}
	}
}
