package main

import (
	"fmt"
	"time"
)

type BloodType int

const (
	O  BloodType = 0
	A  BloodType = 1
	B  BloodType = 2
	AB BloodType = 3
)

type Person struct {
	FirstName string
	LastName  string
	BloodTYpe BloodType
	Birthday  string
}

var people = []*Person{
	{"山田", "太郎", 0, "1960-01-02"},  // 60
	{"鈴木", "一郎", A, "1970-03-04"},  // 50
	{"佐藤", "真一", A, "1980-05-06"},  // 40
	{"加藤", "一茶", B, "1990-07-08"},  // 30
	{"伊藤", "浩史", AB, "2000-09-10"}, // 20
}

func main() {
	fmt.Printf("全体平均年齢: %.02f 歳\n", 0.0)
	fmt.Printf("O型平均年齢: %.02f 歳\n", 0.0)
	fmt.Printf("A型平均年齢: %.02f 歳\n", 0.0)
	fmt.Printf("B型平均年齢: %.02f 歳\n", 0.0)
	fmt.Printf("AB型平均年齢: %.02f 歳\n", 0.0)
}
