package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Birthday  string
}

var people = []*Person{
	{"山田", "太郎", "1960-01-02"},
	{"鈴木", "一郎", "1970-03-04"},
	{"佐藤", "真一", "1980-05-06"},
	{"加藤", "一茶", "1990-07-08"},
	{"伊藤", "浩史", "2000-09-10"},
}

func main() {
	fmt.Printf("平均年齢: X 歳\n")
}
