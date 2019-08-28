package endpass

import "net/http"

func (ts *TestSuite) TestScopes() {
	ts.testServer = createServer(
		http.StatusOK, MIMEApplicationJSONCharsetUTF8, map[string]interface{}{
			"scopes": []string{
				"wallet:accounts:read",
				"user:email:read",
			},
		},
	)
	ts.client.baseUrl = ts.testServer.URL

	scopes, err := ts.client.Scopes()
	ts.NoError(err)
	ts.NotEmpty(scopes)
	ts.Contains(scopes, "user:email:read")
	ts.Len(scopes, 2)
}
