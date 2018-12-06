package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	err := doIt()
	if err != nil {
		fmt.Println("bad", err)
	}

	fmt.Println("Done")
}

func doIt() error {
	f, err := os.Open("in.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var countThree int
	var countTwo int

	for scanner.Scan() {
		hasThree, hasTwo := countLetters(scanner.Text())

		countThree += hasThree
		countTwo += hasTwo
	}

	fmt.Println("answer:", countThree*countTwo)

	return nil
}

func countLetters(s string) (int, int) {
	letterCounts := make(map[rune]int)

	for _, l := range s {
		letterCounts[l]++
	}

	var hasThree, hasTwo int

	for _, c := range letterCounts {

		if c == 3 && hasThree == 0 {
			hasThree = 1
		}

		if c == 2 && hasTwo == 0 {
			hasTwo = 1
		}
	}

	return hasThree, hasTwo
}
