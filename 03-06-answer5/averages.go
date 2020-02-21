package main

import (
	"time"
)

func averages(people []*Person) (float32, map[BloodType]float32) {
	now := time.Now()
	ages := make([]int, len(people))
	agesMap := map[BloodType][]int{
		O:  []int{},
		A:  []int{},
		B:  []int{},
		AB: []int{},
	}
	for idx, person := range people {
		d, err := time.Parse("2006-01-02", person.Birthday)
		if err != nil {
			panic(err)
		}
		age := now.Year() - d.Year()
		ages[idx] = age
		agesMap[person.BloodType] = append(agesMap[person.BloodType], age)
	}

	sum := 0
	for _, age := range ages {
		sum += age
	}

	sumMap := map[BloodType]int{}
	for bloodType, ages := range agesMap {
		s := 0
		for _, age := range ages {
			s += age
		}
		sumMap[bloodType] = s
	}

	avgMap := map[BloodType]float32{
		O:  float32(sumMap[O]) / float32(len(agesMap[O])),
		A:  float32(sumMap[A]) / float32(len(agesMap[A])),
		B:  float32(sumMap[B]) / float32(len(agesMap[B])),
		AB: float32(sumMap[AB]) / float32(len(agesMap[AB])),
	}

	return float32(sum) / float32(len(people)), avgMap
}
