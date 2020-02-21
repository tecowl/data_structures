package main

import (
	"testing"
)

func TestPersonAge(t *testing.T) {
	patterns := []struct {
		person *Person
		age    int
	}{
		{&Person{"山本", "磯六", B, "1950-01-02"}, 70},
		{&Person{"山田", "太郎", O, "1960-01-02"}, 60},
		{&Person{"鈴木", "一郎", A, "1970-03-04"}, 50},
		{&Person{"佐藤", "真一", A, "1980-05-06"}, 40},
		{&Person{"加藤", "一茶", B, "1990-07-08"}, 30},
		{&Person{"伊藤", "浩史", AB, "2000-09-10"}, 20},
	}

	for _, ptn := range patterns {
		if ptn.age != ptn.person.Age() {
			t.Fatalf("The age of %v must be %d but was %d", *ptn.person, ptn.age, ptn.person.Age())
		}
	}
}
