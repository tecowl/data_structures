package main

type People []*Person

func (s People) Ages() Ints {
	r := make(Ints, len(s))
	for idx, i := range s {
		r[idx] = i.Age()
	}
	return r
}
