package smetrics

import (
	"testing"
)

type jarocase struct {
	s string
	t string
	r float64
}

type levenshteincase struct {
	s     string
	t     string
	icost int
	dcost int
	scost int
	r     int
}

type soundexcase struct {
	s string
	t string
}

func TestJaroWinkler(t *testing.T) {
	cases := []jarocase{
		{"AL", "AL", 1.0},
		{"MARTHA", "MARHTA", 0.9611111111111111},
		{"JONES", "JOHNSON", 0.8323809523809523},
		{"ABCVWXYZ", "CABVWXYZ", 0.9625},
		{"A", "B", 0},
		{"ABCDEF", "123456", 0},
		{"AAAAAAAAABCCCC", "AAAAAAAAABCCCC", 1},
	}

	for _, c := range cases {
		if r := JaroWinkler(c.s, c.t, 0.7, 4); r != c.r {
			t.Errorf("%v, instead of %v", r, c.r)
		}
	}
}
