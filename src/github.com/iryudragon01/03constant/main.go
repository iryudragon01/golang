package main

import "fmt"

func main() {
	fmt.Println("test")
	makefunction()
}

func arraytest() {
	grades := [3]int{97, 84, 88}
	var subject [3]string
	subject[0] = "english"
	subject[1] = "physics"
	subject[2] = "chemical"

	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Subject: %v\n", subject)

	subject = [...]string{"a", "b", "c"}
	fmt.Println(subject)

	var twoDarray [3][3]int
	twoDarray[0] = [3]int{1, 2, 3}
	twoDarray[1] = [3]int{3, 4, 5}
	twoDarray[2] = [3]int{6, 7, 8}

	fmt.Printf("\n2D array : %v\n", twoDarray)

	// pointer array
	a := [...]int{1, 2, 3}
	b := &a
	b[1] = 5 //change value both a and b variable
	fmt.Println(a)
	fmt.Println(b)

}

func slicearray() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a
	b[1] = 5    //change value both a and b variable
	c := a[:]   //slice all element
	d := a[3:]  //slice from 4th to the end
	e := a[:6]  //slice first 6th elements
	f := a[3:6] //slice 4th , 5th and 6th elements
	f[2] = 11	// change all value from a[5] - f[2] 
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
