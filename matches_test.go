package smetrics

import (
	"flag"
	"log"
	"testing"
)

var run *bool

func init() {
	run = flag.Bool("test.matches", false, "get some match comparison output")
}

func TestMatches(t *testing.T) {

	if !*run {
		t.Skip()
	}

	matches := []struct {
		A string
		B string
	}{

		{"Manchester United", "Manchester United"},
		{"Manchester United", "Man Utd"},
		{"Manchester United", "Man United"},
		{"Manchester United", "Newcastle United"},
		{"Manchester City", "Manchester United"},
		{"Man Utd", "Man City"},
		{"Man Utd", "Manchester City"},
	}

	for _, m := range matches {
		log.Printf("Ukkonen for %v: %v", m, Ukkonen(m.A, m.B, 1, 1, 2))
		log.Printf("WagnerFischer for %v: %v", m, WagnerFischer(m.A, m.B, 1, 1, 2))
		log.Printf("Jaro for %v: %v", m, Jaro(m.A, m.B))
		log.Printf("Jaro-Winkler for %v: %v", m, JaroWinkler(m.A, m.B, 0.7, 5))
	}

}
