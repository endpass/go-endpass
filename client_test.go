package endpass

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite

	// test client
	c *Client

	// test server
	srv *httptest.Server
}

func (ts *TestSuite) SetupSuite() {
	ts.srv = httptest.NewServer(ts)
	ts.c = NewClient("abc123", "def456")
	ts.c.baseUrl = ts.srv.URL
}

// handler for API requests
func (ts *TestSuite) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fp := fmt.Sprintf("./testdata%s.json", r.URL.Path)
	body, err := ioutil.ReadFile(fp)
	ts.NoError(err)
	w.Write(body)
}

func TestClient(t *testing.T) {
	ts := &TestSuite{}
	suite.Run(t, ts)
}
