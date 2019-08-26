package endpass

type User struct {
	ID     string   `json:"id"`
	Email  string   `json:"email"`
	Phones []*Phone `json:"phones"`
}

type UserAddress struct {
	ApartmentNumber string  `json:"apartmentNumber"`
	StreetNumber    string  `json:"streetNumber"`
	Street          string  `json:"street"`
	City            string  `json:"city"`
	StateRegion     string  `json:"stateRegion"`
	Country         string  `json:"country"`
	PostalCode      string  `json:"postalCode"`
	Lat             float64 `json:"lat"`
	Lng             float64 `json:"lng"`
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

func (c *Client) UserAddress() (*UserAddress, error) {
	resp := &UserAddress{}
	r, err := c.Get("/user/address")
	if err != nil {
		return nil, err
	}
	err = c.parseResponse(r, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
