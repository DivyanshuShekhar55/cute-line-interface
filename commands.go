package main

func (t *Terminal) getAllCommands() map[string]cmd{
		cmds := map[string]cmd{
		"help": {
			name:     "help",
			desc:     "show a list of cmd and their usage",
			callback: func(){ help(t)},
		},
		"exit": {
			name:     "exit",
			desc:     "exit the terminal also ctrl+c works",
			callback: exit,
		},
		"users": {
			name:     "users",
			desc:     "get users from json placeholder",
			callback: func() { ViewUserList(t.HttpClient) },
		},
		"monkey": {
			name: "monkey",
			desc:"show user details",
			callback: func(){},
		},
	}

	return cmds

}

func (t *Terminal) getCommand(cmd_name string) cmd {

	cmds := t.getAllCommands()
	return cmds[cmd_name]
}
