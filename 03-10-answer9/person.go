package main

import (
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	BloodType BloodType
	Birthday  string
}

func (x *Person) Age() int {
	now := time.Now()
	d, err := time.Parse("2006-01-02", x.Birthday)
	if err != nil {
		panic(err)
	}
	return now.Year() - d.Year()
}
