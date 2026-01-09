package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	c := newHttpClient()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Cute Line Interface")

	for {
		scanner.Scan()
		text := scanner.Text()

		out := tokenise(text)
		cmd := getCommand(out[0], &c)
		cmd.callback()

		fmt.Println()
	}

}
