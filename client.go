package endpass

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

type Client struct {
	*http.Client

	baseUrl string
}

// TODO TokenSource implementation
func NewClient() *Client {
	var src oauth2.TokenSource
	c := oauth2.NewClient(context.Background(), src)
	return &Client{
		Client:  c,
		baseUrl: APIBase,
	}
}
