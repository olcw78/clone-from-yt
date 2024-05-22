package src

import "fmt"

type Person struct {
	name          string
	age           int
	favouriteFood []string
}

func makePerson() Person {
	return Person{
		name: "yoon",
		age:  28,
		favouriteFood: []string{
			"pizza", "pasta", "jang-jo-rim",
		},
	}
}

func RunStruct() {
	yoon := makePerson()
	fmt.Println(yoon)
}
