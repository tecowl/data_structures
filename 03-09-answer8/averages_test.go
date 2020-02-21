package main

import (
	"testing"
)

func TestAverages(t *testing.T) {
	{
		allAvg, avgMap := averages([]*Person{
			{"山本", "磯六", B, "1950-01-02"},  // 70
			{"山田", "太郎", O, "1960-01-02"},  // 60
			{"鈴木", "一郎", A, "1970-03-04"},  // 50
			{"佐藤", "真一", A, "1980-05-06"},  // 40
			{"加藤", "一茶", B, "1990-07-08"},  // 30
			{"伊藤", "浩史", AB, "2000-09-10"}, // 20
		})
		if 45 != allAvg {
			t.Fatalf("allAvg must be 45 but was %f", allAvg)
		}
		expectedMap := map[BloodType]float32{
			O:  60.0,
			A:  45.0,
			B:  50.0,
			AB: 20.0,
		}
		for bloodType, expected := range expectedMap {
			if expected != avgMap[bloodType] {
				t.Fatalf("allAvg must be %f but was %f", expected, avgMap[bloodType])
			}
		}
	}

}
