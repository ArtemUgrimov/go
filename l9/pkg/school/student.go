package school

import "math/rand"

type Student struct {
	Id    int
	Marks map[string]int
}

func NewStudent(id int) Student {
	s := Student{
		Id: id,
	}
	s.Marks = make(map[string]int)

	s.Marks["math"] = rand.Int() % 100
	s.Marks["philosophy"] = rand.Int() % 100
	s.Marks["computer science"] = 80 + rand.Int()%20
	s.Marks["physics"] = rand.Int() % 100

	return s
}
