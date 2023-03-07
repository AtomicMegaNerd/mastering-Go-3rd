package main

import "fmt"

func main() {
	s1 := []int64{1, 2, 3, 4, 5}
	fmt.Printf("s1: len %d, cap %d\n", len(s1), cap(s1))
	s1 = append(s1, 6)
	fmt.Printf("s1: len %d, cap %d\n", len(s1), cap(s1))

	// I can specify a capacity
	s2 := make([]int64, 5, 100)
	fmt.Printf("s2: len %d, cap %d\n", len(s2), cap(s2))
	s2[0] = 6
	s2[1] = 7
	s2[2] = 8
	s2[3] = 9
	s2[4] = 10

	// ... explodes our slice into each element which is then passed to
	// append. This is how we append two slices together.
	s1 = append(s1, s2...)
	fmt.Println(s1)

}
