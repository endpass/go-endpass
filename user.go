package endpass

type User struct {
	ID     string   `json:"id"`
	Email  string   `json:"email"`
	Phones []*Phone `json:"phones"`
}

type Phone struct {
	ID        string `json:"id"`
	CreatedAt int64  `json:"createdAt"`
	Status    string `json:"status"`
	Country   string `json:"country"`
	Number    string `json:"number"`
}

func (c *Client) User() (*User, error) {
	resp := &User{}
	r, err := c.Get("/user")
	if err != nil {
		return nil, err
	}
	err = c.parseResponse(r, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
