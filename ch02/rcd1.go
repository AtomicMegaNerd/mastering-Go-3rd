package main

import (
	"fmt"
	"time"
)

func main() {

	nowUnix := time.Now().Unix()
	fmt.Printf("%d\n", nowUnix)
}
