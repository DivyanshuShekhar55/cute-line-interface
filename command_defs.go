package main

import (
	"cute-line-interface/httpx"
	"cute-line-interface/list"
	"cute-line-interface/monkey"
	"cute-line-interface/utils"
	"encoding/json"
	"fmt"

	"net/http"
	"os"
)

func help(terminal *Terminal) {
	cmds := terminal.getAllCommands()
	for _, value := range cmds {
		fmt.Printf("\033[92m\n%s> \033[38;2;175;76;171m%s\033[0m\n", value.name, value.desc)
		fmt.Print(utils.AddDivider("white", 80))
	}

}

func exit() {
	os.Exit(0)
}

func getUser(c *httpx.Client) func() []_User {


	return func() []_User {
		req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/users", nil)
		if err != nil {

			utils.LogError(err)
			return nil
		}

		res, err := c.HttpClient.Do(req)
		if err != nil {
			utils.LogError(err)
			return nil
		}

		defer res.Body.Close()

		var users []User
		if err := json.NewDecoder(res.Body).Decode(&users); err != nil {
			utils.LogError(err)
			return nil
		}

		var result []_User
		for _, u := range users {

			u:= _User{
				name: u.Name,
				email: u.Email,
				username: u.Username,
				website: u.Website,
			}
			result = append(result, u)
		}

		return result
	}

}

func ViewUserList(c *httpx.Client) {
	users := getUser(c)()
	names := []string{}
	for _, item := range users{
		names=append(names, item.name)
	}
	list.List(names)

}

func PrintTable(c *httpx.Client) {
	data  := getUser(c)()
	table := monkey.NewTable()
	table = table.Header([]string{"User", "Email", "Usernmae", "Website"})
	for _, u:= range data {
		table= table.Row([]string{u.name, u.email, u.username, u.website})
	}
	table.Render("magentaBright", "cyan")
}
