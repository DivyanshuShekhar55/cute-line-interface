package main

import (

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
	fmt.Println("\033[92m\nexit> \033[38;2;175;76;171mTexits the cmd terminal\033[0m")
	fmt.Println(addDivider("white"))
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
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
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
			fmt.Println("error occured", err)
			return
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			fmt.Println("error occured", nil)
			return
		}

		defer res.Body.Close()

		

	}
}
