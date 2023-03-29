package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	f, err := os.Open("/dev/random")
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var seed int64
	binary.Read(f, binary.LittleEndian, &seed)
	fmt.Printf("Generated seed %d\n", seed)

	rand.Seed(seed)

	i := rand.Int()
	fmt.Printf("Generated random number %d\n", i)
}
