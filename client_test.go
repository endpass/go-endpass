package endpass

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/oauth2"
)

type TestSuite struct {
	suite.Suite
	client     *Client
	testServer *httptest.Server
}

func (ts *TestSuite) SetupSuite() {
	ts.client = NewClient("clientID", "clientSecret", []string{"1111"}, "12345", "12345")
	ts.client.token = &oauth2.Token{}
	ts.client.baseClient = &http.Client{}
	ts.client.clientWithTokenSource = &http.Client{}
}

func (ts *TestSuite) AfterTest(suiteName, testName string) {
	ts.testServer.Close()
}

func (ts *TestSuite) SetBaseUrl(url string) {
	ts.client.baseUrl = url
}

func TestClient(t *testing.T) {
	ts := &TestSuite{}
	suite.Run(t, ts)
}
