package main

import (
	"bufio"
	"cute-line-interface/httpx"
	"fmt"
	"os"
)

func main() {

	c := httpx.NewHttpClient()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\033[92mWelcome to\033[1;5;96m Cute Line Interface\033[0m")
	fmt.Println("\033[92mType Help For Commands \033[5;92m...\033[0m")

	for {
		scanner.Scan()
		text := scanner.Text()

		out := tokenise(text)
		cmd := getCommand(out[0], &c)
		cmd.callback()

		fmt.Println()
	}

}
