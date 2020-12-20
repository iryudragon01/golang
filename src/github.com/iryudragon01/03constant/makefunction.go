package main

import "fmt"

func makefunction() {
	a := make([]int, 3, 10)
	fmt.Printf("len a :=%v value =%v\n", len(a), a)
	fmt.Printf("capacity := %v\n", cap(a))

	b := []int{}
	fmt.Printf("len b :=%v value =%v\n", len(b), b)
	fmt.Printf("capacity := %v\n", cap(b))

	b = append(b, 5)
	fmt.Printf("len b :=%v value =%v\n", len(b), b)
	fmt.Printf("capacity := %v\n", cap(b))
	b = append(b, 6, 7, 8, 9)
	fmt.Printf("len b :=%v value =%v\n", len(b), b)
	fmt.Printf("capacity := %v\n", cap(b))

	//concatenate slice
	c := []int{}
	c = append(c, 1)
	c = append(c, []int{2, 3, 4}...)
	c = append(c, b...)
	fmt.Printf("slice c := %v\n", c)

	d := append(c[1:3], c[5:len(c)-1]...)
	fmt.Printf("d := %v\n", d)

	d[0] = 99
	d = append(d, 77)
	fmt.Printf("c =: %v\n", c)
	fmt.Printf("d =: %v\n", d)
}
