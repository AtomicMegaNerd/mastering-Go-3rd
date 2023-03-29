package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func lineByLine(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			return err
		}
		fmt.Println(line)
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error, missing argument")
		os.Exit(1)
	}

	file := os.Args[1]
	err := lineByLine(file)

	if err != nil {
		fmt.Printf("Error reading file %s", file)
		os.Exit(1)
	}
}
