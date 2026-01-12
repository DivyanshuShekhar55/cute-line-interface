package main

import "cute-line-interface/httpx"

func getCommand(cmd_name string, c *httpx.Client) cmd {

	cmds := map[string]cmd{
		"help": {
			name:     "help",
			desc:     "show a list of cmd and their usage",
			callback: help,
		},
		"exit": {
			name:     "exit",
			desc:     "exit the terminal also ctrl+c works",
			callback: exit,
		},
		"users": {
			name:     "users",
			desc:     "get users from json placeholder",
			callback: func() { ViewUserList(c) },
		},
	}

	return cmds[cmd_name]
}
