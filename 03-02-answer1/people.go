package main

type People []*Person

func (s People) Ages() Ints {
	r := make(Ints, len(s))
	for idx, i := range s {
		r[idx] = i.Age()
	}
	return r
}

func (s People) Select(f func(*Person) bool) People {
	r := People{}
	for _, i := range s {
		if f(i) {
			r = append(r, i)
		}
	}
	return r
}

func (s People) SelectByBloodType(t BloodType) People {
	return s.Select(func(i *Person) bool {
		return i.BloodType == t
	})
}
