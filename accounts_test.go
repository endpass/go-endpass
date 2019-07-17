package endpass

func (ts *TestSuite) TestAccounts() {
	accounts, err := ts.c.Accounts()
	ts.NoError(err)
	ts.NotEmpty(accounts)
	ts.Contains(accounts, "0x123")
	ts.Len(accounts, 2)
}

func (ts *TestSuite) TestActiveAccount() {
	account, err := ts.c.ActiveAccount()
	ts.NoError(err)
	ts.NotEmpty(account)
	ts.Equal("0x123", account)
}
