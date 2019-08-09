package endpass

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

type Client struct {
	baseClient            *http.Client
	clientWithTokenSource *http.Client
	oauth2Config          *oauth2.Config
	token                 *oauth2.Token
	baseUrl               string
	state                 string
}

func NewClient(
	clientId, clientSecret string, scopes []string, state string, redirectURL string,
) *Client {
	config := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		RedirectURL:  redirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  OAuth2BaseURL + "/api/v1.1/oauth/auth",
			TokenURL: OAuth2BaseURL + "/api/v1.1/oauth/token",
		},
	}
	return &Client{
		baseClient:   defaultHttpClient(2 * time.Second),
		baseUrl:      PublicAPIBaseURL,
		oauth2Config: config,
		state:        state,
	}
}

func (c *Client) HttpClient() *http.Client {
	return c.baseClient
}

func (c *Client) SetHttpClient(httpClient *http.Client) {
	c.baseClient = httpClient
	c.clientWithTokenSource = nil
}

func (c *Client) Token() *oauth2.Token {
	return c.token
}

func (c *Client) SetToken(token *oauth2.Token) {
	c.token = token
	c.clientWithTokenSource = nil
}

func (c *Client) AuthCodeURL() string {
	return c.oauth2Config.AuthCodeURL(c.state)
}

func (c *Client) IsStateValid(state string) bool {
	return c.state == state
}

func (c *Client) Exchange(code string) error {
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, c.baseClient)
	var err error
	c.token, err = c.oauth2Config.Exchange(ctx, code)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Get(path string) (*http.Response, error) {
	if c.token == nil {
		return nil, ErrNoAccessToken
	}
	if c.clientWithTokenSource == nil {
		ctx := context.WithValue(context.Background(), oauth2.HTTPClient, c.baseClient)
		c.clientWithTokenSource = c.oauth2Config.Client(ctx, c.token)
	}
	reqUrl := c.baseUrl + path
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.clientWithTokenSource.Do(req)
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
	return json.Unmarshal(body, v)
}

// check200Response converts response codes not equal 200 to errors
func check200Response(resp *http.Response) (*http.Response, error) {
	if resp.StatusCode == http.StatusOK {
		return resp, nil
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var bodyString string
	if utf8.Valid(bodyBytes) {
		bodyString = string(bodyBytes)
	} else {
		bodyString = "<binary response>"
	}
	return nil, NewErrorHTTPResponse(resp.Status, bodyString)
}
