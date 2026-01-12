package httpx 

import (
	"net/http"
	"time"
)

type Client struct {
	HttpClient *http.Client
}

func NewHttpClient() Client {
	return Client{
		HttpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

