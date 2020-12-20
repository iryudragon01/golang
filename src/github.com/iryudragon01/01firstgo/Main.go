package main

import "fmt"

var (
	author string = "Bandahn Namboonma"
	date   string = "19 december 2020"
)

func main() {
	var i int
	i = 5
	var j int = 6
	k := 7

	var bu bool = false

	fmt.Println("hello go")
	fmt.Printf("%v,%T\n", i, i)
	fmt.Println(j, k)
	fmt.Printf("%v,%T", bu, bu)
	fmt.Println("\n", author)
}
