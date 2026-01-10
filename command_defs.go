package main

import (
	"encoding/json"
	"fmt"

	"net/http"
	"os"
)

type cmd struct {
	name     string
	desc     string
	callback func()
}

func help() {
	fmt.Println("\033[92m\nexit> \033[38;2;175;76;171mexits the cmd terminal\033[0m")
	fmt.Println(addDivider("white", 80))
}

func exit() {
	os.Exit(0)
}

func getUser(c *client) func() {

	type Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	}

	type Geo struct {
		Lat string `json:"lat"`
		Lng string `json:"lng"`
	}

	type Address struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     Geo    `json:"geo"`
	}

	type User struct {
		ID       int     `json:"id"`
		Name     string  `json:"name"`
		Username string  `json:"username"`
		Email    string  `json:"email"`
		Address  Address `json:"address"`
		Phone    string  `json:"phone"`
		Website  string  `json:"website"`
		Company  Company `json:"company"`
	}

	return func() {
		req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/users", nil)
		if err != nil {

			return
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			logError(err)
			return
		}

		defer res.Body.Close()

		var users []User
		if err := json.NewDecoder(res.Body).Decode(&users); err != nil {
			logError(err)
			return
		}

		for _, u := range users {

			user_str := u.Name + " " + u.Website + " " + u.Email
			user_item := turnText(user_str, "magenta", false,
				false)
			fmt.Println(user_item)
		}

	}
}
