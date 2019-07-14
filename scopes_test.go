package endpass

func (ts *TestSuite) TestScopes() {
	scopes, err := ts.c.Scopes()
	ts.NoError(err)
	ts.NotEmpty(scopes)
	ts.Contains(scopes, "user:email:read")
}
