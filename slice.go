package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	//               0, 1, 2, 3,  4,  5

	var s []int = primes[1:4]
	// a[low:high] // low included high excluded
	fmt.Println(s)
}
