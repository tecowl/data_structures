package main

type Ints []int

func (s Ints) Sum() int {
	r := 0
	for _, i := range s {
		r += i
	}
	return r
}

func (s Ints) Average() float32 {
	return float32(s.Sum()) / float32(len(s))
}
