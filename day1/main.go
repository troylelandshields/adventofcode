package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	err := doIt(0, make(map[int]bool))
	if err != nil {
		fmt.Println("bad", err)
	}

	fmt.Println("Done")
}

func doIt(startValue int, frequencies map[int]bool) error {

	fmt.Println("current", startValue)

	f, err := os.Open("in.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	current := startValue

	for scanner.Scan() {
		_, exists := frequencies[current]
		if exists {
			fmt.Println("answer:", current)
			return nil
		}

		frequencies[current] = true

		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}

		current += num
	}

	return doIt(current, frequencies)
}
