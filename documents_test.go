package endpass

import (
	"io/ioutil"
)

func (ts *TestSuite) TestDocuments() {
	documents, err := ts.c.Documents()
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
	document, err := ts.c.Document("1")
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
	documentFile, err := ts.c.DocumentFile("1")
	ts.NoError(err)
	ts.NotEmpty(documentFile)
	defer documentFile.Close()
	fileBody, err := ioutil.ReadAll(documentFile)
	ts.NoError(err)
	ts.NotEmpty(fileBody)
	ts.Equal("{}\n", string(fileBody))
}

func (ts *TestSuite) TestDocumentFrontFile() {
	documentFile, err := ts.c.DocumentFrontFile("1")
	ts.NoError(err)
	ts.NotEmpty(documentFile)
	defer documentFile.Close()
	fileBody, err := ioutil.ReadAll(documentFile)
	ts.NoError(err)
	ts.NotEmpty(fileBody)
	ts.Equal("{}\n", string(fileBody))
}

func (ts *TestSuite) TestDocumentBackFile() {
	documentFile, err := ts.c.DocumentBackFile("1")
	ts.NoError(err)
	ts.NotEmpty(documentFile)
	defer documentFile.Close()
	fileBody, err := ioutil.ReadAll(documentFile)
	ts.NoError(err)
	ts.NotEmpty(fileBody)
	ts.Equal("{}\n", string(fileBody))
}
