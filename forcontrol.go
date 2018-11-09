package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	person := make([]Person, 2)
	fmt.Println(AverageAge(person))
	person1 := []Person{{"aa", 2},{"bb", 2}}
	fmt.Println(AverageAge(person1))
	fmt.Println(person1[0].ageSum())
}

func AverageAge(person []Person) int {

	if len(person) == 0 {
		return 0
	}
	var count, sum int
	for _, p := range person {
		sum += p.age
		count += 1
	}
	return sum / count
}

func (p Person) ageSum() int{
	return p.age + 1;
}