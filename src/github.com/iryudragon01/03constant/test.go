package main

import "fmt"

const (
	errorSpecialist = iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func test() {
	const myCon int = 42

	fmt.Printf("%v,%T\n", myCon, myCon)
	fmt.Printf("%v\n", errorSpecialist)
	fmt.Printf("%v\n", catSpecialist)
	fmt.Printf("%v\n", dogSpecialist)
	fmt.Printf("%v\n", snakeSpecialist)
}
