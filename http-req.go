package main

import (
	"net/http"
	"time"
)

type client struct {
	httpClient *http.Client
}

func newHttpClient() client {
	return client{
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

