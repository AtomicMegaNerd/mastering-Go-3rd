package main

import "fmt"

func printMessages(messages ...string) {
	for _, message := range messages {
		fmt.Println(message)
	}
}

func main() {
	printMessages("Learning", "Go", "Is", "So", "Much", "Fun")
	messages2 := []string{"Learning", "Go", "Is", "Crazy", "Fun"}

	printMessages(messages2...)
}
