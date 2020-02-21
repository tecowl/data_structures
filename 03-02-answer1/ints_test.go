package main

import (
	"testing"
)

func TestIntsAverage(t *testing.T) {
	patterns := []struct {
		ints Ints
		val  float32
	}{
		{Ints{1, 2, 3}, 2},
		{Ints{4, 5, 6}, 5},
		{Ints{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5.5},
	}

	for _, ptn := range patterns {
		if ptn.val != ptn.ints.Average() {
			t.Fatalf("The average of %v must be %f but was %f", ptn.ints, ptn.val, ptn.ints.Average())
		}
	}
}
