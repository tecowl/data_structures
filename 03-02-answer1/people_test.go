package main

import (
	"testing"
)

func TestPeopleAges(t *testing.T) {
	people := People{
		{"山本", "磯六", B, "1950-01-02"},
		{"山田", "太郎", O, "1960-01-02"},
		{"鈴木", "一郎", A, "1970-03-04"},
		{"佐藤", "真一", A, "1980-05-06"},
		{"加藤", "一茶", B, "1990-07-08"},
		{"伊藤", "浩史", AB, "2000-09-10"},
	}

	expectedAges := []int{70, 60, 50, 40, 30, 20}
	actualAges := people.Ages()

	if len(expectedAges) != len(actualAges) {
		t.Fatalf("The length of ages must be %d but was %d", len(expectedAges), len(actualAges))
	}

	for idx, expectedAge := range expectedAges {
		if expectedAge != actualAges[idx] {
			t.Fatalf("The ages[%d]  must be %d but was %d", idx, expectedAge, actualAges[idx])
		}
	}
}
