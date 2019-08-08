package endpass

func (ts *TestSuite) TestUser() {
	user, err := ts.c.User()
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
	userAddress, err := ts.c.UserAddress()
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
