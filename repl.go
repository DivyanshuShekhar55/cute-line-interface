package main

import (
	"strings"
)

func tokenise(text string) (output []string) {
	lowered_text := strings.ToLower(text)
	output = strings.Fields(lowered_text)
	return output
}


