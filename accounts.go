package endpass

type activeAccountResp struct {
	Address string `json:"address"`
}

func (c *Client) Accounts() ([]string, error) {
	resp := make([]string, 0, 10)
	r, err := c.Get("/accounts")
	if err != nil {
		return nil, err
	}
	err = c.parseResponse(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) ActiveAccount() (string, error) {
	resp := activeAccountResp{}
	r, err := c.Get("/accounts/active")
	if err != nil {
		return "", err
	}
	err = c.parseResponse(r, &resp)
	if err != nil {
		return "", err
	}
	return resp.Address, nil
}
