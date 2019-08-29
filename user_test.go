package endpass

import "net/http"

func (ts *TestSuite) TestUser() {
	ts.testServer = createServer(
		http.StatusOK, MIMEApplicationJSONCharsetUTF8, map[string]interface{}{
			"id":    "0d8c5fa3-c8a5-4c5f-8435-f35aef353f30",
			"email": "user@endpass.com",
			"phones": []interface{}{
				map[string]interface{}{
					"id":        "c4d4ef1c-0d73-4a6f-aad9-7600a8fc79b8",
					"createdAt": 1557220652,
					"status":    "Verified",
					"country":   "7",
					"number":    "7771112233",
				},
			},
		},
	)
	ts.client.baseUrl = ts.testServer.URL

	user, err := ts.client.User()
	ts.NoError(err)
	ts.NotEmpty(user)
	ts.Equal("0d8c5fa3-c8a5-4c5f-8435-f35aef353f30", user.ID)
	ts.Equal("user@endpass.com", user.Email)
	ts.Len(user.Phones, 1)
	ts.Equal("c4d4ef1c-0d73-4a6f-aad9-7600a8fc79b8", user.Phones[0].ID)
	ts.Equal(int64(1557220652), user.Phones[0].CreatedAt)
	ts.Equal("Verified", user.Phones[0].Status)
	ts.Equal("7", user.Phones[0].Country)
	ts.Equal("7771112233", user.Phones[0].Number)
}

func (ts *TestSuite) TestUserAddress() {
	ts.testServer = createServer(
		http.StatusOK, MIMEApplicationJSONCharsetUTF8, map[string]interface{}{
			"apartmentNumber": "apartment number 1",
			"streetNumber":    "street number 2",
			"street":          "street 8",
			"city":            "city 17",
			"stateRegion":     "state region",
			"country":         "big country",
			"postalCode":      "postal code",
			"lat":             1.1,
			"lng":             2.2,
		},
	)
	ts.client.baseUrl = ts.testServer.URL

	userAddress, err := ts.client.UserAddress()
	ts.NoError(err)
	ts.NotEmpty(userAddress)
	ts.NotNil(userAddress)
	ts.Equal("apartment number 1", userAddress.ApartmentNumber)
	ts.Equal("street number 2", userAddress.StreetNumber)
	ts.Equal("street 8", userAddress.Street)
	ts.Equal("city 17", userAddress.City)
	ts.Equal("state region", userAddress.StateRegion)
	ts.Equal("big country", userAddress.Country)
	ts.Equal("postal code", userAddress.PostalCode)
	ts.Equal(1.1, userAddress.Lat)
	ts.Equal(2.2, userAddress.Lng)
}
