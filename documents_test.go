package endpass

import (
	"io/ioutil"
	"net/http"
)

func (ts *TestSuite) TestDocuments() {
	ts.testServer = createServer(
		http.StatusOK, MIMEApplicationJSONCharsetUTF8, []interface{}{
			map[string]interface{}{
				"id":               "f1f80c7a-57c7-4259-8e80-d1bb09932e82",
				"createdAt":        1543336121,
				"status":           "New",
				"documentType":     "Passport",
				"description":      "Custom document description",
				"firstName":        "Test",
				"lastName":         "User",
				"number":           "47313892501",
				"dateOfBirth":      1543323121,
				"dateOfIssue":      1543336122,
				"dateOfExpiry":     1543336123,
				"issuingCountry":   "Country",
				"issuingAuthority": "Authority",
				"issuingPlace":     "Place",
				"address":          "Address",
			},
		},
	)
	ts.client.baseUrl = ts.testServer.URL

	documents, err := ts.client.Documents()
	ts.NoError(err)
	ts.NotEmpty(documents)
	ts.Len(documents, 1)
	document := documents[0]
	ts.Equal("f1f80c7a-57c7-4259-8e80-d1bb09932e82", document.ID)
	ts.Equal(int64(1543336121), document.CreatedAt)
	ts.Equal("New", document.Status)
	ts.Equal("Passport", document.DocumentType)
	ts.Equal("Custom document description", document.Description)
	ts.Equal("Test", document.FirstName)
	ts.Equal("User", document.LastName)
	ts.Equal("47313892501", document.Number)
	ts.Equal(int64(1543323121), document.DateOfBirth)
	ts.Equal(int64(1543336122), document.DateOfIssue)
	ts.Equal(int64(1543336123), document.DateOfExpiry)
	ts.Equal("Country", document.IssuingCountry)
	ts.Equal("Authority", document.IssuingAuthority)
	ts.Equal("Place", document.IssuingPlace)
	ts.Equal("Address", document.Address)
}

func (ts *TestSuite) TestDocument() {
	ts.testServer = createServer(
		http.StatusOK, MIMEApplicationJSONCharsetUTF8, map[string]interface{}{
			"id":               "f1f80c7a-57c7-4259-8e80-d1bb09932e82",
			"createdAt":        1543336121,
			"status":           "New",
			"documentType":     "Passport",
			"description":      "Custom document description",
			"firstName":        "Test",
			"lastName":         "User",
			"number":           "47313892501",
			"dateOfBirth":      1543323121,
			"dateOfIssue":      1543336122,
			"dateOfExpiry":     1543336123,
			"issuingCountry":   "Country",
			"issuingAuthority": "Authority",
			"issuingPlace":     "Place",
			"address":          "Address",
		},
	)
	ts.client.baseUrl = ts.testServer.URL

	document, err := ts.client.Document("1")
	ts.NoError(err)
	ts.NotEmpty(document)
	ts.Equal("f1f80c7a-57c7-4259-8e80-d1bb09932e82", document.ID)
	ts.Equal(int64(1543336121), document.CreatedAt)
	ts.Equal("New", document.Status)
	ts.Equal("Passport", document.DocumentType)
	ts.Equal("Custom document description", document.Description)
	ts.Equal("Test", document.FirstName)
	ts.Equal("User", document.LastName)
	ts.Equal("47313892501", document.Number)
	ts.Equal(int64(1543323121), document.DateOfBirth)
	ts.Equal(int64(1543336122), document.DateOfIssue)
	ts.Equal(int64(1543336123), document.DateOfExpiry)
	ts.Equal("Country", document.IssuingCountry)
	ts.Equal("Authority", document.IssuingAuthority)
	ts.Equal("Place", document.IssuingPlace)
	ts.Equal("Address", document.Address)
}

func (ts *TestSuite) TestDocumentFile() {
	ts.testServer = createServer(
		http.StatusOK, MIMEApplicationJSONCharsetUTF8, map[string]interface{}{},
	)
	ts.client.baseUrl = ts.testServer.URL

	documentFile, err := ts.client.DocumentFile("1")
	ts.NoError(err)
	ts.NotEmpty(documentFile)
	defer documentFile.Close()
	fileBody, err := ioutil.ReadAll(documentFile)
	ts.NoError(err)
	ts.NotEmpty(fileBody)
	ts.Equal("{}", string(fileBody))
}

func (ts *TestSuite) TestDocumentFrontFile200() {
	ts.testServer = createServer(
		http.StatusOK, MIMEApplicationJSONCharsetUTF8, map[string]interface{}{},
	)
	ts.client.baseUrl = ts.testServer.URL

	documentFile, err := ts.client.DocumentFrontFile("1")
	ts.NoError(err)
	ts.NotEmpty(documentFile)
	defer documentFile.Close()
	fileBody, err := ioutil.ReadAll(documentFile)
	ts.NoError(err)
	ts.NotEmpty(fileBody)
	ts.Equal("{}", string(fileBody))
}

func (ts *TestSuite) TestDocumentFrontFile204() {
	ts.testServer = createServer(
		http.StatusNoContent, MIMEApplicationJSONCharsetUTF8, map[string]interface{}{},
	)
	ts.client.baseUrl = ts.testServer.URL

	documentFile, err := ts.client.DocumentFrontFile("1")
	ts.Empty(documentFile)
	ts.IsType(ErrFileNotUploaded, err)
}

func (ts *TestSuite) TestDocumentBackFile200() {
	ts.testServer = createServer(
		http.StatusOK, MIMEApplicationJSONCharsetUTF8, map[string]interface{}{},
	)
	ts.client.baseUrl = ts.testServer.URL

	documentFile, err := ts.client.DocumentBackFile("1")
	ts.NoError(err)
	ts.NotEmpty(documentFile)
	defer documentFile.Close()
	fileBody, err := ioutil.ReadAll(documentFile)
	ts.NoError(err)
	ts.NotEmpty(fileBody)
	ts.Equal("{}", string(fileBody))
}

func (ts *TestSuite) TestDocumentBackFile204() {
	ts.testServer = createServer(
		http.StatusNoContent, MIMEApplicationJSONCharsetUTF8, map[string]interface{}{},
	)
	ts.client.baseUrl = ts.testServer.URL

	documentFile, err := ts.client.DocumentBackFile("1")
	ts.Empty(documentFile)
	ts.IsType(ErrFileNotUploaded, err)
}
