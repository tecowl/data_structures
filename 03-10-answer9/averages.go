package main

func averages(people People) (float32, map[BloodType]float32) {
	return people.Ages().Average(),
		map[BloodType]float32{
			O:  people.SelectByBloodType(O).Ages().Average(),
			A:  people.SelectByBloodType(A).Ages().Average(),
			B:  people.SelectByBloodType(B).Ages().Average(),
			AB: people.SelectByBloodType(AB).Ages().Average(),
		}
}
