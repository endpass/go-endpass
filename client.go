package endpass

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/net/proxy"
	"golang.org/x/oauth2"
)

type Client struct {
	httpClient   *http.Client
	oauth2Config *oauth2.Config
	dialer       proxy.Dialer
	baseUrl      string
	state        string
}

func NewClient(
	clientId string, scopes []string, state string, redirectURL string,
) *Client {
	config := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: "",
		Scopes:       scopes,
		RedirectURL:  redirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  OAuth2BaseURL + "/api/v1.1/oauth/auth",
			TokenURL: OAuth2BaseURL + "/api/v1.1/oauth/token",
		},
	}
	return &Client{
		httpClient:   nil,
		baseUrl:      PublicAPIBaseURL,
		oauth2Config: config,
		dialer:       nil,
		state:        state,
	}
}

// SetDialer install custom Dialer. Needed for example for working through proxy server.
func (c *Client) SetDialer(dialer proxy.Dialer) {
	c.dialer = dialer
}

func (c *Client) AuthCodeURL() string {
	return c.oauth2Config.AuthCodeURL(c.state)
}

func (c *Client) IsStateValid(state string) bool {
	return c.state == state
}

func (c *Client) Exchange(code string) error {
	ctx := context.Background()

	httpTransport := &http.Transport{}
	if c.dialer != nil {
		httpTransport.Dial = c.dialer.Dial
	}
	httpClient := &http.Client{
		Timeout:   2 * time.Second,
		Transport: httpTransport,
	}

	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	token, err := c.oauth2Config.Exchange(ctx, code)
	if err != nil {
		return err
	}

	c.httpClient = c.oauth2Config.Client(ctx, token)

	return nil
}

func (c *Client) Get(path string) (*http.Response, error) {
	if c.httpClient == nil {
		return nil, ErrNoAccessToken
	}
	reqUrl := c.baseUrl + path
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return check200Response(resp)
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

// check200Response converts response codes not equal 200 to errors
func check200Response(resp *http.Response) (*http.Response, error) {
	if resp.StatusCode == http.StatusOK {
		return resp, nil
	}
	return nil, NewErrorHTTPResponse(resp)
}
