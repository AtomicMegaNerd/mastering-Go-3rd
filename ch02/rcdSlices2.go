package main

import "fmt"

func main() {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s1)

	l1 := len(s1)

	// First 5 elements, both the below are equivalent
	fmt.Println(s1[0:5])
	fmt.Println(s1[:5])

	// Last two elements
	fmt.Println(s1[l1-2 : l1])
	fmt.Println(s1[l1-2:])

	// First 5 elements with a capacity of 10
	t := s1[:5:10]
	fmt.Println(t, len(t), cap(t))

	// In a slice [x:y:z] the capacity is z - x
	t = s1[2:5:10]
	fmt.Println(t, len(t), cap(t))

	// Last 2 elements with a capacity of 10-8 = 2
	// In a slice [x:y:z] the capacity is z - x
	t = s1[l1-2 : l1 : 10]
	fmt.Println(t, len(t), cap(t))

}
