package main

import (
	"bufio"
	"cute-line-interface/httpx"
	"fmt"
	"os"
)

func main() {

	client := httpx.NewHttpClient()
	terminal := Terminal{HttpClient: &client}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\033[92mWelcome to\033[1;5;96m Cute Line Interface\033[0m")
	fmt.Println("\033[92mType Help For Commands \033[5;92m...\033[0m")

	for {
		scanner.Scan()
		text := scanner.Text()

		out := tokenise(text)
		cmd := terminal.getCommand(out[0])
		cmd.callback()

		fmt.Println()
	}

}
