package main

import "fmt"

const (
	_  = iota             // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota) // 2^10
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func iodatest() {
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)
	fmt.Println(KB)

}
