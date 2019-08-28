package endpass

import "net/http"

func (ts *TestSuite) TestAccounts() {
	ts.testServer = createServer(
		http.StatusOK, MIMEApplicationJSONCharsetUTF8, []string{"0x123", "0x456"},
	)
	ts.client.baseUrl = ts.testServer.URL

	accounts, err := ts.client.Accounts()
	ts.NoError(err)
	ts.NotEmpty(accounts)
	ts.Contains(accounts, "0x123")
	ts.Len(accounts, 2)
}

func (ts *TestSuite) TestActiveAccount() {
	ts.testServer = createServer(
		http.StatusOK, MIMEApplicationJSONCharsetUTF8, map[string]interface{}{
			"address": "0x123",
		},
	)
	ts.client.baseUrl = ts.testServer.URL

	account, err := ts.client.ActiveAccount()
	ts.NoError(err)
	ts.NotEmpty(account)
	ts.Equal("0x123", account)
}
