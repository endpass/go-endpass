package endpass

type scopesResp struct {
	Scopes []string `json:"scopes"`
}

// Scopes gets the list of scopes your application is auhorized to access
func (c *Client) Scopes() ([]string, error) {
	resp := &scopesResp{}
	r, err := c.Get("/scopes")
	if err != nil {
		return nil, err
	}
	err = c.parseResponse(r, resp)
	if err != nil {
		return nil, err
	}
	return resp.Scopes, nil
}
