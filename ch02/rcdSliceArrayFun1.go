package main

import "fmt"

func change(s []string) {
	s[0] = "Change_function"
}

func main() {

	// This is an array
	a := [4]string{"Zero", "One", "Two", "Three"}
	fmt.Println()
	fmt.Println("a:", a)

	// This slice is connected to the array beneath it. Modifying the slice modifies
	// the array underneath
	var s0 = a[0:1]
	fmt.Println()
	fmt.Println("s0:", s0)
	s0[0] = "s0"

	// This slice is also connected to the array beneath it. Modifying the slice modifies
	// the array underneath
	s1 := a[1:3]
	fmt.Println()
	fmt.Println("s1:", s1)
	s1[0] = "s1_0"
	s1[1] = "s2_0"

	// Show the modifications to the array
	fmt.Println()
	fmt.Println("a:", a)

	fmt.Println()
	fmt.Println("cap of s0:", cap(s0), ",len of s0:", len(s0), ",s0:", s0)
	s0 = append(s0, "n1")
	s0 = append(s0, "n2")
	s0 = append(s0, "n3")
	a[0] = "-N1" // Still connected because cap has not changed
	fmt.Println()
	fmt.Println("cap of s0:", cap(s0), "len of s0:", len(s0), "s0:", s0)
	fmt.Println("cap of a:", cap(a), "len of a:", len(a), "a:", a)

	a[0] = "-n1-"
	a[1] = "-n2-"

	// This breaks the link between s0 and a because the capacity of the slice has increased
	s0 = append(s0, "n4")
	fmt.Println()
	fmt.Println("cap of s0:", cap(s0), "len of s0:", len(s0), "s0:", s0)
	fmt.Println("cap of s1:", cap(s1), "len of s1:", len(s1), "s1:", s1)
	fmt.Println("cap of a:", cap(a), "len of a:", len(a), "a:", a)
}
