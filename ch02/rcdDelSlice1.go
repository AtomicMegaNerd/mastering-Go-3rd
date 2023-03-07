package main

import (
	"errors"
	"fmt"
)

// This function should take an element in a slice and the slice
// and it should remove the element from the slice assuming it exists
// in the slice. If the element is not present an error will be
// returned.
func delElemFromSlice[T comparable](elem T, slice []T) ([]T, error) {
	ix, err := func(elem T, slice []T) (int, error) {
		for ix, e := range slice {
			if e == elem {
				return ix, nil
			}
		}
		return -1, errors.New("element not found")
	}(elem, slice)
	slice = append(slice[:ix], slice[ix+1:]...)
	return slice, err
}

func main() {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s1, len(s1), cap(s1))

	s1, err := delElemFromSlice(5, s1)
	if err != nil {
		panic(err)
	}
	fmt.Println(s1, len(s1), cap(s1))

	s1, err = delElemFromSlice(9, s1)
	if err != nil {
		panic(err)
	}
	fmt.Println(s1, len(s1), cap(s1))
}
