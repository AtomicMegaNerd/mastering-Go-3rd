package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
)

const NUM_ITEMS = 200

type DataHolder struct {
	items [NUM_ITEMS]int
}

func MutateDataHolder(ch chan<- error, dh *DataHolder, ix int) {
	num := rand.Intn(10 - 0)
	if num == 0 {
		ch <- errors.New("zero returned")
	} else {
		dh.items[ix] = num
		ch <- nil
	}
}

func main() {

	wg := &sync.WaitGroup{}
	ch := make(chan error, NUM_ITEMS)
	dh := &DataHolder{}

	for ix := 0; ix < NUM_ITEMS; ix++ {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			MutateDataHolder(ch, dh, ix)
		}(ix)
	}
	wg.Wait()

	for ix := 0; ix < NUM_ITEMS; ix++ {
		err := <-ch
		if err != nil {
			fmt.Printf("Error %s\n", err)
		} else {
			fmt.Printf("items[%d] = %d\n", ix, dh.items[ix])
		}
	}

	close(ch)
}
