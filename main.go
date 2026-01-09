package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Cute Line Interface")

	for {
		scanner.Scan()
		text := scanner.Text()

		out := tokenise(text)
		fmt.Print(">")

		for i := 0; i < len(out); i++ {
			fmt.Print(out[i], " ")
		}
		fmt.Println()
	}

}
