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

	seen := make(map[string]bool)

	for scanner.Scan() {
		s := scanner.Text()

		possible := make(map[string]bool)

		for i := 0; i < len(s); i++ {
			sWithoutLetter := s[0:i] + s[i+1:len(s)]

			if _, ok := seen[sWithoutLetter]; ok {
				fmt.Println("answer:", sWithoutLetter)
				return nil
			}

			possible[sWithoutLetter] = true
		}

		for p := range possible {
			seen[p] = true
		}
	}

	return nil
}
