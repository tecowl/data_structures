package main

type People []*Person

func (s People) Ages() []int {
	r := make([]int, len(s))
	for idx, i := range s {
		r[idx] = i.Age()
	}
	return r
}
