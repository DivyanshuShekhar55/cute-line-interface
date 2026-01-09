package main

import (
	"fmt"
	"os"
)

type cmd struct {
	name     string
	desc     string
	callback func()
}

func help() {
	fmt.Println("\nexit> exits the cmd terminal")
	fmt.Println("---------------------------------------")
}

func exit() {
	os.Exit(0)
}

func getUser(c *client) func() {
	return func() {
		
		
		
	}
}
