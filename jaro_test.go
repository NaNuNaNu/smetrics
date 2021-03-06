package smetrics

import "testing"

func TestJaro(t *testing.T) {
	cases := []jarocase{
		{"AL", "AL", 1.0},
		{"MARTHA", "MARHTA", 0.9444444444444445},
		{"JONES", "JOHNSON", 0.7904761904761904},
		{"ABCVWXYZ", "CABVWXYZ", 0.9583333333333334},
		{"A", "B", 0},
		{"ABCDEF", "123456", 0},
		{"AAAAAAAAABCCCC", "AAAAAAAAABCCCC", 1},
	}

	for _, c := range cases {
		if r := Jaro(c.s, c.t); r != c.r {
			t.Errorf("%v, instead of %v", r, c.r)
		}
	}
}
