package endpass

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)

type Client struct {
	*http.Client

	baseUrl string
}

// TODO TokenSource implementation
func NewClient(clientId, clientSecret string) *Client {
	var src oauth2.TokenSource
	c := oauth2.NewClient(context.Background(), src)
	return &Client{
		Client:  c,
		baseUrl: APIBase,
	}
}

func (c *Client) Get(path string) (*http.Response, error) {
	reqUrl := c.baseUrl + path
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// parses response as JSON into an object
// v is a pointer to an object that can be unmarshalled into
func (c *Client) parseResponse(r *http.Response, v interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	err = json.Unmarshal(body, v)
	return err
}
