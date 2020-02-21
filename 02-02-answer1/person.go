package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Birthday  string
}

var people = []*Person{
	{"山田", "太郎", "1960-01-02"}, // 60
	{"鈴木", "一郎", "1970-03-04"}, // 50
	{"佐藤", "真一", "1980-05-06"}, // 40
	{"加藤", "一茶", "1990-07-08"}, // 30
	{"伊藤", "浩史", "2000-09-10"}, // 20
}

func main() {
	now := time.Now()
	ages := make([]int, len(people))
	for idx, person := range people {
		d, err := time.Parse("2006-01-02", person.Birthday)
		if err != nil {
			panic(err)
		}
		ages[idx] = now.Year() - d.Year()
	}

	sum := 0
	for _, age := range ages {
		sum += age
	}

	fmt.Printf("平均年齢: %.02f 歳\n", float32(sum)/float32(len(people)))
}
